//ff:func feature=seg type=helper control=sequence
//ff:what Buffer.Size — memory the rows and their sort index take
package seg

func (b *Buffer) Size() int64 {
	return int64(len(b.data)) + 4*int64(b.Rows())
}
