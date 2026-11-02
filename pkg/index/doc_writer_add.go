//ff:func feature=index type=writer control=sequence
//ff:what docWriter.add — one entry
package index

func (w *docWriter) add(e DocEntry) error {
	var b [docSize]byte
	w.sum.Write(e.append(b[:0]))
	w.count++
	_, err := w.bw.Write(b[:])
	return err
}
