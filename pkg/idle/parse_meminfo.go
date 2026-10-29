//ff:func feature=idle type=decoder control=iteration dimension=1
//ff:what parseMeminfo — MemTotal and MemAvailable of /proc/meminfo in bytes (the file gives kB)
package idle

import (
	"fmt"
	"strconv"
	"strings"
)

// parseMeminfo reads "MemTotal:  16000000 kB" lines.
func parseMeminfo(b []byte) (uint64, uint64, error) {
	var total, avail uint64
	for _, line := range strings.Split(string(b), "\n") {
		f := strings.Fields(line)
		if len(f) < 2 || (f[0] != "MemTotal:" && f[0] != "MemAvailable:") {
			continue
		}
		v, err := strconv.ParseUint(f[1], 10, 64)
		if err != nil {
			return 0, 0, err
		}
		if f[0] == "MemTotal:" {
			total = v * 1024
		} else {
			avail = v * 1024
		}
	}
	if total == 0 {
		return 0, 0, fmt.Errorf("%w: no MemTotal", ErrUnsupported)
	}
	return total, avail, nil
}
