//ff:func feature=seg type=rule control=iteration dimension=1
//ff:what File.checkOrder — rows of a region do not descend (whole rows when sorted, else the partition slot of other rows)
package seg

import (
	"bytes"
	"fmt"
)

// checkOrder streams g with a cursor and compares neighbours: their first n
// bytes must not descend, nor whole rows when sorted.
func (f *File) checkOrder(g Group, n int, sorted bool) error {
	c := &cursor{g: g}
	var prev []byte
	for {
		row, ok, err := c.next()
		if !ok {
			return err
		}
		if prev != nil && (bytes.Compare(prev[:n], row[:n]) > 0 || (sorted && bytes.Compare(prev, row) > 0)) {
			return fmt.Errorf("%w: %s: rows out of order", ErrFormat, f.Path)
		}
		prev = append(prev[:0], row...)
	}
}
