//ff:func feature=sig type=helper control=sequence
//ff:what Schema.Words — uint64 words per signature (a 16- or 32-bit signature still takes one word)
package sig

// Words is ceil(Width/64).
func (s Schema) Words() int {
	return (s.Width + 63) / 64
}
