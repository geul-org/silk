//ff:func feature=index type=helper control=sequence
//ff:what Bitmap.Put — set or clear a bit, growing when i is beyond N
package index

func (b *Bitmap) Put(i uint32, v bool) {
	if i >= b.N {
		b.N = i + 1
		b.B = append(b.B, make([]byte, int((uint64(b.N)+7)/8)-len(b.B))...)
	}
	if v {
		b.B[i/8] |= 1 << (i % 8)
		return
	}
	b.B[i/8] &^= 1 << (i % 8)
}
