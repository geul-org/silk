//ff:func feature=seg type=helper control=sequence
//ff:what File.otherRegion — the rows of the other partition (after every partition row)
package seg

// otherRegion locates the other partition.
func (f *File) otherRegion() region {
	r, ro := f.Info.RowWidth()
	return region{off: int64(f.Info.Rows) * int64(r), n: f.Info.Other, w: ro}
}
