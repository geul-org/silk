//ff:func feature=sig type=rule control=iteration dimension=1
//ff:what Pattern.Exact — every term holds through its exact codes, so the row answers without a document
package sig

func (p Pattern) Exact(bits []byte) bool {
	for _, t := range p.Terms {
		if !allows(t.Exact, get(bits, t.Off, t.Width)) {
			return false
		}
	}
	return true
}
