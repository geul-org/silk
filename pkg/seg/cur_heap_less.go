//ff:func feature=seg type=helper control=sequence
//ff:what curHeap.Less — heap.Interface: by key, then row
package seg

import "bytes"

func (h curHeap) Less(i, j int) bool {
	if c := bytes.Compare(h[i].key, h[j].key); c != 0 {
		return c < 0
	}
	return bytes.Compare(h[i].row(), h[j].row()) < 0
}
