//ff:func feature=seg type=builder control=sequence
//ff:what Buffer.Free — forget the rows and release the memory
package seg

func (b *Buffer) Free() {
	b.data, b.idx = nil, nil
}
