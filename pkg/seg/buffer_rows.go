//ff:func feature=seg type=helper control=sequence
//ff:what Buffer.Rows — the rows held
package seg

func (b *Buffer) Rows() int {
	return len(b.data) / (b.pw + b.rw)
}
