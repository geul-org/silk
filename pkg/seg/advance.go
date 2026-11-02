//ff:func feature=seg type=scanner control=sequence
//ff:what advance — step the top cursor and restore the heap
package seg

import "container/heap"

func advance(h *curHeap, c *cursor) {
	if c.next() {
		heap.Fix(h, 0)
		return
	}
	heap.Pop(h)
}
