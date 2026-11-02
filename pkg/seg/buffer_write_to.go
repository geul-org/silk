//ff:func feature=seg type=writer control=iteration dimension=1
//ff:what Buffer.WriteTo — sort and write every row
package seg

func (b *Buffer) WriteTo(w *Writer) error {
	b.sort()
	for _, i := range b.idx {
		r := b.at(i)
		if err := w.Add(r[:b.pw], r[b.pw:]); err != nil {
			return err
		}
	}
	return nil
}
