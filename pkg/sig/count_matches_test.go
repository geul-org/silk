//ff:func feature=sig type=rule control=iteration dimension=1
//ff:what countMatches — rows a pattern matches, read one by one
package sig

func countMatches(p Pattern, rows [][]byte) int {
	n := 0
	for _, r := range rows {
		if p.Match(r) {
			n++
		}
	}
	return n
}
