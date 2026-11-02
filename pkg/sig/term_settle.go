//ff:func feature=sig type=scanner control=sequence
//ff:what Term.settle — bring term i of bits up to an allowed code; the term to look at next (0 after a carry); false when the bits overflow
package sig

func (t Term) settle(b []byte, i int) (int, bool) {
	c := get(b, t.Off, t.Width)
	n, ok := nextCode(t.Spans, c)
	if !ok {
		return 0, carry(b, t.Off)
	}
	if n != c {
		put(b, t.Off, t.Width, n)
		clearFrom(b, t.Off+t.Width)
	}
	return i + 1, true
}
