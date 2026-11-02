//ff:type feature=seg type=model
//ff:what Writer — one segment being written in order: rows, then the partition table spilled to a side file, then the footer
package seg

import (
	"bufio"
	"hash"
	"os"
)

type Writer struct {
	info  Info
	path  string
	f, tf *os.File
	bw    *bufio.Writer
	tw    *bufio.Writer
	sum   hash.Hash
	part  []byte
	last  []byte
	open  bool
	start uint64
	rows  uint64
	parts uint64
}
