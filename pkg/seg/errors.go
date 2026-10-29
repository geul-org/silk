package seg

import "errors"

var (
	// ErrFormat means segment bytes break the format (docs/segments.md).
	ErrFormat = errors.New("silk/seg: bad segment file")
	// ErrOrder means a writer got partitions out of key order.
	ErrOrder = errors.New("silk/seg: partitions out of order")
)
