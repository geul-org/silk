//ff:func feature=index type=helper control=iteration dimension=1
//ff:what Bitmap.Count — the number of live sequence numbers
package index

import "math/bits"

// Count is the population count.
func (b Bitmap) Count() int {
	n := 0
	for _, x := range b.B {
		n += bits.OnesCount8(x)
	}
	return n
}
