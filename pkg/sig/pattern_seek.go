//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what Pattern.seek — the smallest bits at or above b that satisfy every term; false when none remain
package sig

func (p Pattern) seek(b []byte) bool {
	for i := 0; i < len(p.Terms); {
		next, ok := p.Terms[i].settle(b, i)
		if !ok {
			return false
		}
		i = next
	}
	return true
}
