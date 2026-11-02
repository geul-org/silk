//ff:func feature=seg type=helper control=sequence
//ff:what curHeap.Pop — heap.Interface
package seg

func (h *curHeap) Pop() any {
	o := *h
	n := len(o)
	x := o[n-1]
	*h = o[:n-1]
	return x
}
