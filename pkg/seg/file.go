//ff:type feature=seg type=model
//ff:what File — an opened segment: its mmap, rows and partition table
package seg

type File struct {
	Info
	Rows  uint64
	Parts uint64
	Path  string
	b     []byte
	table []byte
}
