//ff:func feature=sig type=rule control=iteration dimension=1
//ff:what Pattern.Match — every term holds on these bits
package sig

func (p Pattern) Match(bits []byte) bool {
	for _, t := range p.Terms {
		if !allows(t.Spans, get(bits, t.Off, t.Width)) {
			return false
		}
	}
	return true
}
