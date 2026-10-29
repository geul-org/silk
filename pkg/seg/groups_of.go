//ff:func feature=seg type=reader control=iteration dimension=1
//ff:what GroupsOf — every partition group of several segments, sorted by key (input order kept for equal keys): a merge's input
package seg

// GroupsOf prepares Merge.
func GroupsOf(fs []*File) ([]Group, error) {
	var gs []Group
	for _, f := range fs {
		g, err := f.Groups()
		if err != nil {
			return nil, err
		}
		gs = append(gs, g...)
	}
	SortGroups(gs)
	return gs, nil
}
