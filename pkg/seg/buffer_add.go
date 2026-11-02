//ff:func feature=seg type=builder control=sequence
//ff:what Buffer.Add — hold one row
package seg

func (b *Buffer) Add(part, row []byte) {
	b.data = append(append(b.data, part...), row...)
}
