//ff:func feature=sig type=helper control=iteration dimension=1
//ff:what clearFrom — zero every bit from offset off
package sig

func clearFrom(b []byte, off int) {
	for i := off; i < len(b)*8; i++ {
		b[i/8] &^= 0x80 >> (i % 8)
	}
}
