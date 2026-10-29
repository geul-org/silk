//ff:func feature=seg type=helper control=iteration dimension=1
//ff:what SortRows — fixed-width rows sorted as byte strings (sort key, document, offset), into a new buffer
package seg

import (
	"bytes"
	"slices"
)

// SortRows sorts b's rows of r bytes.
func SortRows(b []byte, r int) []byte {
	views := make([][]byte, 0, len(b)/r)
	for ; len(b) >= r; b = b[r:] {
		views = append(views, b[:r])
	}
	slices.SortFunc(views, bytes.Compare)
	return bytes.Join(views, nil)
}
