//ff:func feature=index type=helper control=iteration dimension=1
//ff:what rowsOf — rows of some segments
package index

func rowsOf(ss []SegRef) uint64 {
	var n uint64
	for _, s := range ss {
		n += s.Rows
	}
	return n
}
