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

package collector

import (
	"container/heap"

	"github.com/edwindvinas/bleve/search"
)

type collectStoreHeap struct {
	heap    search.DocumentMatchCollection
	compare collectorCompare
}

func newStoreHeap(cap int, compare collectorCompare) *collectStoreHeap {
	rv := &collectStoreHeap{
		heap:    make(search.DocumentMatchCollection, 0, cap),
		compare: compare,
	}
	heap.Init(rv)
	return rv
}

func (c *collectStoreHeap) Add(doc *search.DocumentMatch) {
	heap.Push(c, doc)
}

func (c *collectStoreHeap) RemoveLast() *search.DocumentMatch {
	return heap.Pop(c).(*search.DocumentMatch)
}

func (c *collectStoreHeap) Final(skip int, fixup collectorFixup) (search.DocumentMatchCollection, error) {
	count := c.Len()
	size := count - skip
	if size <= 0 {
		return make(search.DocumentMatchCollection, 0), nil
	}
	rv := make(search.DocumentMatchCollection, size)
	for count > 0 {
		count--

		if count >= skip {
			size--
			doc := heap.Pop(c).(*search.DocumentMatch)
			rv[size] = doc
			err := fixup(doc)
			if err != nil {
				return nil, err
			}
		}
	}
	return rv, nil
}

// heap interface implementation

func (c *collectStoreHeap) Len() int {
	return len(c.heap)
}

func (c *collectStoreHeap) Less(i, j int) bool {
	so := c.compare(c.heap[i], c.heap[j])
	return -so < 0
}

func (c *collectStoreHeap) Swap(i, j int) {
	c.heap[i], c.heap[j] = c.heap[j], c.heap[i]
}

func (c *collectStoreHeap) Push(x interface{}) {
	c.heap = append(c.heap, x.(*search.DocumentMatch))
}

func (c *collectStoreHeap) Pop() interface{} {
	var rv *search.DocumentMatch
	rv, c.heap = c.heap[len(c.heap)-1], c.heap[:len(c.heap)-1]
	return rv
}
