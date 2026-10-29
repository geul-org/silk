//ff:func feature=sig type=encoder control=iteration dimension=1
//ff:what putBits — OR the low width bits of v into a signature at bit offset bit (may cross a word boundary)
package sig

// putBits writes v into s[bit : bit+width).
func putBits(s []uint64, bit, width int, v uint64) {
	for i := 0; i < width; i++ {
		s[(bit+i)/64] |= (v >> i & 1) << ((bit + i) % 64)
	}
}
