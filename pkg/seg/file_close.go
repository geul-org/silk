//ff:func feature=seg type=reader control=sequence
//ff:what File.Close — release the segment file
package seg

// Close closes the file.
func (f *File) Close() error {
	return f.c.Close()
}
