//ff:func feature=seg type=helper control=sequence
//ff:what Buffer.at — the bytes of row i (partition first)
package seg

func (b *Buffer) at(i uint32) []byte {
	w := b.pw + b.rw
	return b.data[int(i)*w : (int(i)+1)*w]
}
