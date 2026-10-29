//ff:func feature=recmap type=reader control=sequence
//ff:what File.Close — release the underlying file (nothing for a File made by Read)
package recmap

// Close closes the file Open opened.
func (f *File) Close() error {
	if f.c == nil {
		return nil
	}
	return f.c.Close()
}
