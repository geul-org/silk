//ff:func feature=seg type=reader control=iteration dimension=1
//ff:what File.Groups — every partition key of a segment as a Group: table partitions, then the other partition split by key
package seg

// Groups lists f's partitions in key order within each region.
func (f *File) Groups() ([]Group, error) {
	out := make([]Group, 0, len(f.Parts))
	for i, p := range f.Parts {
		out = append(out, Group{Key: p.Key, f: f, rg: f.partRegion(i)})
	}
	return f.otherGroups(out)
}
