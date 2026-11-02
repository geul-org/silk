//ff:func feature=sig type=helper control=iteration dimension=1
//ff:what carry — add one at the last bit before off and zero the rest; false when the bits overflow
package sig

func carry(b []byte, off int) bool {
	clearFrom(b, off)
	for i := off - 1; i >= 0; i-- {
		m := byte(0x80) >> (i % 8)
		b[i/8] ^= m
		if b[i/8]&m != 0 {
			return true
		}
	}
	return false
}
