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

package query

import (
	"regexp"
	"strings"

	"github.com/edwindvinas/bleve/index"
	"github.com/edwindvinas/bleve/mapping"
	"github.com/edwindvinas/bleve/search"
	"github.com/edwindvinas/bleve/search/searcher"
)

type RegexpQuery struct {
	Regexp   string `json:"regexp"`
	FieldVal string `json:"field,omitempty"`
	BoostVal *Boost `json:"boost,omitempty"`
	compiled *regexp.Regexp
}

// NewRegexpQuery creates a new Query which finds
// documents containing terms that match the
// specified regular expression.  The regexp pattern
// SHOULD NOT include ^ or $ modifiers, the search
// will only match entire terms even without them.
func NewRegexpQuery(regexp string) *RegexpQuery {
	return &RegexpQuery{
		Regexp: regexp,
	}
}

func (q *RegexpQuery) SetBoost(b float64) {
	boost := Boost(b)
	q.BoostVal = &boost
}

func (q *RegexpQuery) Boost() float64 {
	return q.BoostVal.Value()
}

func (q *RegexpQuery) SetField(f string) {
	q.FieldVal = f
}

func (q *RegexpQuery) Field() string {
	return q.FieldVal
}

func (q *RegexpQuery) Searcher(i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	field := q.FieldVal
	if q.FieldVal == "" {
		field = m.DefaultSearchField()
	}
	err := q.compile()
	if err != nil {
		return nil, err
	}

	return searcher.NewRegexpSearcher(i, q.compiled, field, q.BoostVal.Value(), options)
}

func (q *RegexpQuery) Validate() error {
	return q.compile()
}

func (q *RegexpQuery) compile() error {
	if q.compiled == nil {
		// require that pattern NOT be anchored to start and end of term
		actualRegexp := q.Regexp
		if strings.HasPrefix(actualRegexp, "^") {
			actualRegexp = actualRegexp[1:] // remove leading ^
		}
		// do not attempt to remove trailing $, it's presence is not
		// known to interfere with LiteralPrefix() the way ^ does
		// and removing $ introduces possible ambiguities with escaped \$, \\$, etc
		var err error
		q.compiled, err = regexp.Compile(actualRegexp)
		if err != nil {
			return err
		}
	}
	return nil
}
