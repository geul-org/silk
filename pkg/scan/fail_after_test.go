//ff:type feature=scan type=model
//ff:what failAfter — a ReaderAt that serves n reads, then fails
package scan

import "io"

// failAfter counts reads.
type failAfter struct {
	r io.ReaderAt
	n int
}
