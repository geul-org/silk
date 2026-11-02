//ff:func feature=sig type=builder control=iteration dimension=1
//ff:what cross — the overlap of two span lists
package sig

func cross(a, b []Span) []Span {
	var out []Span
	for _, x := range a {
		out = append(out, overlaps(x, b)...)
	}
	return out
}
