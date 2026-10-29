//ff:func feature=seg type=helper control=iteration dimension=1
//ff:what allSorted — every group comes from a sorted run
package seg

// allSorted reports whether a streaming merge keeps order.
func allSorted(gs []Group) bool {
	for _, g := range gs {
		if !g.f.Info.Sorted {
			return false
		}
	}
	return true
}
