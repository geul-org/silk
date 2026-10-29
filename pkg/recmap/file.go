//ff:type feature=recmap type=model
//ff:what File — an opened map file: header, key ranges and documents in memory, off[] read on demand with ReadAt
package recmap

import "io"

// File reads one map file. off[] stays on disk; Offsets reads one key's
// range. It is safe for concurrent reads.
type File struct {
	Header Header
	Keys   []KeyRange
	Docs   []DocEntry
	r      io.ReaderAt
	c      io.Closer
}
