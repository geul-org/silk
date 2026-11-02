//ff:func feature=index type=helper control=iteration dimension=1
//ff:what addSegs — add segment files to keep
package index

func addSegs(keep map[string]bool, ss []SegRef) {
	for _, s := range ss {
		keep[s.File] = true
	}
}
