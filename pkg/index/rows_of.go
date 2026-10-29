//ff:func feature=index type=helper control=iteration dimension=1
//ff:what rowsOf — total rows of segment references
package index

// rowsOf sums Rows.
func rowsOf(rs []SegRef) uint64 {
	n := uint64(0)
	for _, r := range rs {
		n += r.Rows
	}
	return n
}
