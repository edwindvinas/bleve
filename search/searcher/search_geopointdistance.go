//  Copyright (c) 2017 Couchbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package searcher

import (
	"github.com/edwindvinas/bleve/geo"
	"github.com/edwindvinas/bleve/index"
	"github.com/edwindvinas/bleve/numeric"
	"github.com/edwindvinas/bleve/search"
)

func NewGeoPointDistanceSearcher(indexReader index.IndexReader, centerLon,
	centerLat, dist float64, field string, boost float64,
	options search.SearcherOptions) (search.Searcher, error) {

	// compute bounding box containing the circle
	topLeftLon, topLeftLat, bottomRightLon, bottomRightLat :=
		geo.ComputeBoundingBox(centerLon, centerLat, dist)

		// build a searcher for the box
	boxSearcher, err := boxSearcher(indexReader,
		topLeftLon, topLeftLat, bottomRightLon, bottomRightLat,
		field, boost, options)
	if err != nil {
		return nil, err
	}

	// wrap it in a filtering searcher which checks the actual distance
	return NewFilteringSearcher(boxSearcher,
		buildDistFilter(indexReader, field, centerLon, centerLat, dist)), nil
}

// boxSearcher builds a searcher for the described bounding box
// if the desired box crosses the dateline, it is automatically split into
// two boxes joined through a disjunction searcher
func boxSearcher(indexReader index.IndexReader,
	topLeftLon, topLeftLat, bottomRightLon, bottomRightLat float64,
	field string, boost float64, options search.SearcherOptions) (
	search.Searcher, error) {
	if bottomRightLon < topLeftLon {
		// cross date line, rewrite as two parts

		leftSearcher, err := NewGeoBoundingBoxSearcher(indexReader,
			-180, bottomRightLat, bottomRightLon, topLeftLat,
			field, boost, options, false)
		if err != nil {
			return nil, err
		}
		rightSearcher, err := NewGeoBoundingBoxSearcher(indexReader,
			topLeftLon, bottomRightLat, 180, topLeftLat, field, boost, options, false)
		if err != nil {
			_ = leftSearcher.Close()
			return nil, err
		}

		boxSearcher, err := NewDisjunctionSearcher(indexReader,
			[]search.Searcher{leftSearcher, rightSearcher}, 0, options)
		if err != nil {
			_ = leftSearcher.Close()
			_ = rightSearcher.Close()
			return nil, err
		}
		return boxSearcher, nil
	}

	// build geoboundinggox searcher for that bounding box
	boxSearcher, err := NewGeoBoundingBoxSearcher(indexReader,
		topLeftLon, bottomRightLat, bottomRightLon, topLeftLat, field, boost,
		options, false)
	if err != nil {
		return nil, err
	}
	return boxSearcher, nil
}

func buildDistFilter(indexReader index.IndexReader, field string,
	centerLon, centerLat, maxDist float64) FilterFunc {
	return func(d *search.DocumentMatch) bool {
		var lon, lat float64
		var found bool
		err := indexReader.DocumentVisitFieldTerms(d.IndexInternalID,
			[]string{field}, func(field string, term []byte) {
				// only consider the values which are shifted 0
				prefixCoded := numeric.PrefixCoded(term)
				shift, err := prefixCoded.Shift()
				if err == nil && shift == 0 {
					i64, err := prefixCoded.Int64()
					if err == nil {
						lon = geo.MortonUnhashLon(uint64(i64))
						lat = geo.MortonUnhashLat(uint64(i64))
						found = true
					}
				}
			})
		if err == nil && found {
			dist := geo.Haversin(lon, lat, centerLon, centerLat)
			if dist <= maxDist/1000 {
				return true
			}
		}
		return false
	}
}
