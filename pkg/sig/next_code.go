//ff:func feature=sig type=helper control=iteration dimension=1
//ff:what nextCode — the smallest allowed code at or above c
package sig

func nextCode(spans []Span, c uint64) (uint64, bool) {
	best, ok := uint64(0), false
	for _, s := range spans {
		v := max(s.Lo, c)
		if v <= s.Hi && (!ok || v < best) {
			best, ok = v, true
		}
	}
	return best, ok
}
