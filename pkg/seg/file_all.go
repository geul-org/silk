//ff:func feature=seg type=helper control=sequence
//ff:what File.All — every row of the segment as one region (partition keys aside)
package seg

func (f *File) All() Region {
	return Region{B: f.b[:f.Rows*uint64(f.RowWidth)], W: f.RowWidth}
}
