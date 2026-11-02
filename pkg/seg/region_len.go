//ff:func feature=seg type=helper control=sequence
//ff:what Region.Len — the number of rows
package seg

func (r Region) Len() int {
	return len(r.B) / r.W
}
