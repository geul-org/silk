//ff:type feature=seg type=model
//ff:what Writer — a segment being written: temporary file, running SHA-256, partition table, other rows, the open partition's buffered rows
package seg

import (
	"bufio"
	"hash"
	"os"
)

// Writer takes partitions in ascending key order (Begin, Add, End) and
// writes their rows; Finish adds the other rows, the table and the footer and
// renames the file into place. A partition stays buffered until it reaches
// Min rows; one that ends below Min goes to the other partition.
type Writer struct {
	Path    string
	Min     int
	info    Info
	f       *os.File
	bw      *bufio.Writer
	h       hash.Hash
	parts   []Part
	other   []byte
	key     []byte
	open    bool
	buf     []byte
	n       uint64
	written int64
}
