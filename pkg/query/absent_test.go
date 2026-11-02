//ff:func feature=query type=rule control=iteration dimension=1
//ff:what absent — the first path missing from either answer set; empty when none is
package query

func absent(want []string, sets ...map[string]bool) string {
	for _, p := range want {
		if !sets[0][p] || !sets[1][p] {
			return p
		}
	}
	return ""
}
