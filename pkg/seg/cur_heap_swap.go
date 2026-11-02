//ff:func feature=seg type=helper control=sequence
//ff:what curHeap.Swap — heap.Interface
package seg

func (h curHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
