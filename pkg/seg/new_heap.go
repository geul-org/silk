//ff:func feature=seg type=builder control=iteration dimension=1
//ff:what newHeap — a heap of the non-empty inputs
package seg

import "container/heap"

func newHeap(in []*File) *curHeap {
	h := &curHeap{}
	for _, f := range in {
		c := &cursor{f: f}
		if c.load() {
			*h = append(*h, c)
		}
	}
	heap.Init(h)
	return h
}
