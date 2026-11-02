//ff:func feature=seg type=builder control=sequence
//ff:what Buffer.Absorb — take another buffer's rows and empty it
package seg

func (b *Buffer) Absorb(o *Buffer) {
	b.data = append(b.data, o.data...)
	o.data = o.data[:0]
}
