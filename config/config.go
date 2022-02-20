//  Copyright (c) 2015 Couchbase, Inc.
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

package config

import (
	// token maps
	_ "github.com/edwindvinas/bleve/analysis/tokenmap"

	// fragment formatters
	_ "github.com/edwindvinas/bleve/search/highlight/format/ansi"
	_ "github.com/edwindvinas/bleve/search/highlight/format/html"

	// fragmenters
	_ "github.com/edwindvinas/bleve/search/highlight/fragmenter/simple"

	// highlighters
	_ "github.com/edwindvinas/bleve/search/highlight/highlighter/ansi"
	_ "github.com/edwindvinas/bleve/search/highlight/highlighter/html"
	_ "github.com/edwindvinas/bleve/search/highlight/highlighter/simple"

	// char filters
	_ "github.com/edwindvinas/bleve/analysis/char/html"
	_ "github.com/edwindvinas/bleve/analysis/char/regexp"
	_ "github.com/edwindvinas/bleve/analysis/char/zerowidthnonjoiner"

	// analyzers
	_ "github.com/edwindvinas/bleve/analysis/analyzer/custom"
	_ "github.com/edwindvinas/bleve/analysis/analyzer/keyword"
	_ "github.com/edwindvinas/bleve/analysis/analyzer/simple"
	_ "github.com/edwindvinas/bleve/analysis/analyzer/standard"
	_ "github.com/edwindvinas/bleve/analysis/analyzer/web"

	// token filters
	_ "github.com/edwindvinas/bleve/analysis/token/apostrophe"
	_ "github.com/edwindvinas/bleve/analysis/token/compound"
	_ "github.com/edwindvinas/bleve/analysis/token/edgengram"
	_ "github.com/edwindvinas/bleve/analysis/token/elision"
	_ "github.com/edwindvinas/bleve/analysis/token/keyword"
	_ "github.com/edwindvinas/bleve/analysis/token/length"
	_ "github.com/edwindvinas/bleve/analysis/token/lowercase"
	_ "github.com/edwindvinas/bleve/analysis/token/ngram"
	_ "github.com/edwindvinas/bleve/analysis/token/shingle"
	_ "github.com/edwindvinas/bleve/analysis/token/stop"
	_ "github.com/edwindvinas/bleve/analysis/token/truncate"
	_ "github.com/edwindvinas/bleve/analysis/token/unicodenorm"

	// tokenizers
	_ "github.com/edwindvinas/bleve/analysis/tokenizer/exception"
	_ "github.com/edwindvinas/bleve/analysis/tokenizer/regexp"
	_ "github.com/edwindvinas/bleve/analysis/tokenizer/single"
	_ "github.com/edwindvinas/bleve/analysis/tokenizer/unicode"
	_ "github.com/edwindvinas/bleve/analysis/tokenizer/web"
	_ "github.com/edwindvinas/bleve/analysis/tokenizer/whitespace"

	// date time parsers
	_ "github.com/edwindvinas/bleve/analysis/datetime/flexible"
	_ "github.com/edwindvinas/bleve/analysis/datetime/optional"

	// languages
	_ "github.com/edwindvinas/bleve/analysis/lang/ar"
	_ "github.com/edwindvinas/bleve/analysis/lang/bg"
	_ "github.com/edwindvinas/bleve/analysis/lang/ca"
	_ "github.com/edwindvinas/bleve/analysis/lang/cjk"
	_ "github.com/edwindvinas/bleve/analysis/lang/ckb"
	_ "github.com/edwindvinas/bleve/analysis/lang/cs"
	_ "github.com/edwindvinas/bleve/analysis/lang/el"
	_ "github.com/edwindvinas/bleve/analysis/lang/en"
	_ "github.com/edwindvinas/bleve/analysis/lang/eu"
	_ "github.com/edwindvinas/bleve/analysis/lang/fa"
	_ "github.com/edwindvinas/bleve/analysis/lang/fr"
	_ "github.com/edwindvinas/bleve/analysis/lang/ga"
	_ "github.com/edwindvinas/bleve/analysis/lang/gl"
	_ "github.com/edwindvinas/bleve/analysis/lang/hi"
	_ "github.com/edwindvinas/bleve/analysis/lang/hy"
	_ "github.com/edwindvinas/bleve/analysis/lang/id"
	_ "github.com/edwindvinas/bleve/analysis/lang/in"
	_ "github.com/edwindvinas/bleve/analysis/lang/it"
	_ "github.com/edwindvinas/bleve/analysis/lang/pt"

	// kv stores
	_ "github.com/edwindvinas/bleve/index/store/boltdb"
	_ "github.com/edwindvinas/bleve/index/store/goleveldb"
	_ "github.com/edwindvinas/bleve/index/store/gtreap"
	_ "github.com/edwindvinas/bleve/index/store/moss"

	// index types
	_ "github.com/edwindvinas/bleve/index/upsidedown"
)
