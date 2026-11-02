//ff:func feature=seg type=helper control=sequence
//ff:what curHeap.Len — heap.Interface
package seg

func (h curHeap) Len() int { return len(h) }
