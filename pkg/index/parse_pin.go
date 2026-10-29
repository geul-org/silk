//ff:func feature=index type=decoder control=sequence
//ff:what parsePin — a pin file name "<generation>-<pid>-<nanotime>" as generation and process ID
package index

import (
	"strconv"
	"strings"
)

// parsePin reads a pin name; false for anything else.
func parsePin(name string) (uint64, int, bool) {
	f := strings.Split(name, "-")
	if len(f) != 3 {
		return 0, 0, false
	}
	g, err1 := strconv.ParseUint(f[0], 10, 64)
	pid, err2 := strconv.Atoi(f[1])
	return g, pid, err1 == nil && err2 == nil
}
