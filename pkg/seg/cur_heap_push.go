//ff:func feature=seg type=helper control=sequence
//ff:what curHeap.Push — heap.Interface
package seg

func (h *curHeap) Push(x any) { *h = append(*h, x.(*cursor)) }
