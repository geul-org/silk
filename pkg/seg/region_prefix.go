//ff:func feature=seg type=helper control=sequence
//ff:what Region.Prefix — the rows starting with prefix
package seg

func (r Region) Prefix(prefix []byte) Region {
	lo := r.Seek(0, prefix)
	hi := r.After(lo, prefix)
	return Region{B: r.B[lo*r.W : hi*r.W], W: r.W}
}
