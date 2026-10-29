//ff:func feature=sig type=matcher control=iteration dimension=1
//ff:what Pattern.match — every word of a signature agrees with Want under Mask
package sig

// match tests one signature.
func (p Pattern) match(s []uint64) bool {
	for i, m := range p.Mask {
		if s[i]&m != p.Want[i] {
			return false
		}
	}
	return true
}
