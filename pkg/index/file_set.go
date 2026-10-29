//ff:func feature=index type=helper control=iteration dimension=1
//ff:what fileSet — the file names of segment references as a set
package index

// fileSet marks each reference's file.
func fileSet(rs []SegRef) map[string]bool {
	out := map[string]bool{}
	for _, r := range rs {
		out[r.File] = true
	}
	return out
}
