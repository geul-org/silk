//ff:func feature=seg type=helper control=sequence
//ff:what File.partRegion — the rows of partition table entry i
package seg

// partRegion locates partition i.
func (f *File) partRegion(i int) region {
	r, _ := f.Info.RowWidth()
	p := f.Parts[i]
	return region{off: int64(p.Start) * int64(r), n: p.Count, w: r}
}
