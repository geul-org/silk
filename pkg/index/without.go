//ff:func feature=index type=helper control=iteration dimension=1
//ff:what without — segment references whose file is not in the gone set
package index

// without filters rs.
func without(rs []SegRef, gone map[string]bool) []SegRef {
	out := []SegRef{}
	for _, r := range rs {
		if !gone[r.File] {
			out = append(out, r)
		}
	}
	return out
}
