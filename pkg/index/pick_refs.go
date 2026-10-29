//ff:func feature=index type=helper control=iteration dimension=1
//ff:what pickRefs — the references among rs whose file is in files, in files' order
package index

// pickRefs matches by file name.
func pickRefs(rs []SegRef, files []string) []SegRef {
	by := map[string]SegRef{}
	for _, r := range rs {
		by[r.File] = r
	}
	var out []SegRef
	for _, f := range files {
		if r, ok := by[f]; ok {
			out = append(out, r)
		}
	}
	return out
}
