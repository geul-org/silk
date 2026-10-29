//ff:type feature=seg type=model
//ff:what failReader — test double: a ReaderAt that fails after n successful reads
package seg

import "io"

// failReader wraps r.
type failReader struct {
	r io.ReaderAt
	n int
}
