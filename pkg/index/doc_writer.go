//ff:type feature=index type=model
//ff:what docWriter — one document table file being written
package index

import (
	"bufio"
	"hash"
	"os"
)

type docWriter struct {
	path  string
	f     *os.File
	bw    *bufio.Writer
	sum   hash.Hash
	first uint32
	count uint32
}
