//ff:func feature=index type=helper control=iteration dimension=1
//ff:what Bitmap.count — live documents
package index

import "math/bits"

func (b Bitmap) count() uint64 {
	var n uint64
	for _, c := range b.B {
		n += uint64(bits.OnesCount8(c))
	}
	return n
}
