//ff:func feature=index type=helper control=iteration dimension=1
//ff:what refFiles — the file names of segment references
package index

// refFiles lists rs' files in order.
func refFiles(rs []SegRef) []string {
	out := make([]string, 0, len(rs))
	for _, r := range rs {
		out = append(out, r.File)
	}
	return out
}
