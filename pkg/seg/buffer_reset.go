//ff:func feature=seg type=builder control=sequence
//ff:what Buffer.Reset — forget the rows, keeping the memory
package seg

func (b *Buffer) Reset() {
	b.data = b.data[:0]
}
