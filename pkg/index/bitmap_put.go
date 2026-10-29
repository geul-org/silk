//ff:func feature=index type=helper control=sequence
//ff:what Bitmap.Put — set or clear a sequence number's bit, growing the bitmap when i is beyond N
package index

// Put sets bit i to v.
func (b *Bitmap) Put(i uint32, v bool) {
	if i >= b.N {
		b.N = i + 1
		b.B = append(b.B, make([]byte, int(b.N+7)/8-len(b.B))...)
	}
	if v {
		b.B[i/8] |= 1 << (i % 8)
	} else {
		b.B[i/8] &^= 1 << (i % 8)
	}
}
