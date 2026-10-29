//ff:func feature=seg type=reader control=sequence
//ff:what readTail — footer and partition table from the end of a segment of size bytes; the sizes must add up exactly
package seg

import (
	"fmt"
	"io"
	"math/bits"
)

// readTail parses the footer, checks the file size and parses the table.
func readTail(f io.ReaderAt, size int64) (*File, error) {
	if size < footerSize {
		return nil, fmt.Errorf("%w: %d bytes", ErrFormat, size)
	}
	fb := make([]byte, footerSize)
	if _, err := f.ReadAt(fb, size-footerSize); err != nil {
		return nil, err
	}
	i, err := parseFooter(fb)
	if err != nil {
		return nil, err
	}
	r, ro := i.RowWidth()
	h1, a := bits.Mul64(i.Rows, uint64(r))
	h2, b := bits.Mul64(i.Other, uint64(ro))
	t := uint64(i.Parts) * uint64(i.PartWidth+tableExtra)
	if h1|h2 != 0 || a+b+t+footerSize != uint64(size) || a+b < a {
		return nil, fmt.Errorf("%w: %d bytes, footer says %d", ErrFormat, size, a+b+t+footerSize)
	}
	tb := make([]byte, t)
	if _, err := f.ReadAt(tb, int64(a+b)); err != nil {
		return nil, err
	}
	parts, err := parseTable(tb, i)
	if err != nil {
		return nil, err
	}
	return &File{Info: i, Parts: parts, r: f}, nil
}
