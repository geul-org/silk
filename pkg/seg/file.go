//ff:type feature=seg type=model
//ff:what File — an opened segment: footer, partition table in memory, rows read on demand with ReadAt
package seg

import "io"

// File reads one segment; safe for concurrent reads.
type File struct {
	Path  string
	Info  Info
	Parts []Part
	r     io.ReaderAt
	c     io.Closer
}
