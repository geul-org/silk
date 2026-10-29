//ff:func feature=sig type=helper control=iteration dimension=1
//ff:what mask — words with bits [bit, bit+width) set
package sig

// mask builds a words-long bit mask.
func mask(bit, width, words int) []uint64 {
	m := make([]uint64, words)
	for i := bit; i < bit+width; i++ {
		m[i/64] |= 1 << (i % 64)
	}
	return m
}
