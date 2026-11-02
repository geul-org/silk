//ff:func feature=sig type=builder control=iteration dimension=1
//ff:what overlaps — one span's overlap with each of a list
package sig

func overlaps(x Span, b []Span) []Span {
	var out []Span
	for _, y := range b {
		if lo, hi := max(x.Lo, y.Lo), min(x.Hi, y.Hi); lo <= hi {
			out = append(out, Span{lo, hi})
		}
	}
	return out
}
