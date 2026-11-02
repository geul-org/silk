//ff:func feature=sig type=rule control=iteration dimension=1
//ff:what Term.allows — a code lies in spans
package sig

func allows(spans []Span, c uint64) bool {
	for _, s := range spans {
		if c >= s.Lo && c <= s.Hi {
			return true
		}
	}
	return false
}
