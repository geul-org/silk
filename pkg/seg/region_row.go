//ff:func feature=seg type=helper control=sequence
//ff:what Region.Row — row i
package seg

func (r Region) Row(i int) []byte {
	return r.B[i*r.W : (i+1)*r.W]
}
