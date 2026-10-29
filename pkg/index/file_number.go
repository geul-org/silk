//ff:func feature=index type=decoder control=sequence
//ff:what fileNumber — the file counter value in a name like seg/000000000000002a.sseg (0 when unreadable)
package index

import (
	"path"
	"strconv"
	"strings"
)

// fileNumber reads the 16 hex digits of an index file name.
func fileNumber(name string) uint64 {
	b := path.Base(name)
	n, _ := strconv.ParseUint(strings.TrimSuffix(b, path.Ext(b)), 16, 64)
	return n
}
