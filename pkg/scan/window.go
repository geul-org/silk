//ff:type feature=scan type=model
//ff:what window — one goroutine's read buffer over a data pack: bytes [lo, lo+len(buf)) and the read size to use
package scan

import "io"

// window caches one ReadAt of the data pack.
type window struct {
	data io.ReaderAt
	size uint64
	min  int
	buf  []byte
	lo   uint64
}
