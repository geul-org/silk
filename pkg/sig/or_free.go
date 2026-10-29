//ff:func feature=sig type=helper control=iteration dimension=1
//ff:what orFree — OR m into used when no bit of m is already used; false (used unchanged) on overlap
package sig

// orFree claims m's bits in used.
func orFree(used, m []uint64) bool {
	for i := range m {
		if used[i]&m[i] != 0 {
			return false
		}
	}
	for i := range m {
		used[i] |= m[i]
	}
	return true
}
