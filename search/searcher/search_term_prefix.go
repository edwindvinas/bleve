//  Copyright (c) 2014 Couchbase, Inc.
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
	"github.com/edwindvinas/bleve/index"
	"github.com/edwindvinas/bleve/search"
)

func NewTermPrefixSearcher(indexReader index.IndexReader, prefix string,
	field string, boost float64, options search.SearcherOptions) (
	search.Searcher, error) {
	// find the terms with this prefix
	fieldDict, err := indexReader.FieldDictPrefix(field, []byte(prefix))
	if err != nil {
		return nil, err
	}

	var terms []string
	tfd, err := fieldDict.Next()
	for err == nil && tfd != nil {
		terms = append(terms, tfd.Term)
		tfd, err = fieldDict.Next()
	}

	return NewMultiTermSearcher(indexReader, terms, field, boost, options)
}
