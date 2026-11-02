//ff:func feature=index type=helper control=iteration dimension=1
//ff:what tierOf — a run's size tier: ⌊log_base rows⌋
package index

func tierOf(rows uint64, base int) int {
	t := 0
	for rows >= uint64(base) {
		rows /= uint64(base)
		t++
	}
	return t
}
