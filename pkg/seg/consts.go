// Package seg is the segment file shared by every SILK section: fixed-width
// rows in byte order, grouped by an optional partition key, opened with mmap.
// Writers stream rows in order; Merge k-way merges sorted segments; Buffer
// sorts rows in memory up to a caller's limit. See docs/segments.md.
package seg

import "errors"

const (
	magic      = "SILKSEG2"
	version    = 1
	footerSize = 128
	bufSize    = 1 << 20
	maxRowSize = 1 << 15
)

var (
	// ErrFormat means a segment file breaks its format.
	ErrFormat = errors.New("silk/seg: bad segment file")
	// ErrOrder means rows were added out of byte order.
	ErrOrder = errors.New("silk/seg: rows out of order")
)
