//ff:func feature=seg type=helper control=sequence
//ff:what File.PartAt — the key and rows of partition i
package seg

func (f *File) PartAt(i uint64) ([]byte, Region) {
	k, s, n := f.entry(i)
	w := uint64(f.RowWidth)
	return k, Region{B: f.b[s*w : (s+n)*w], W: f.RowWidth}
}
