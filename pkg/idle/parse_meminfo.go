//ff:func feature=idle type=decoder control=iteration dimension=1
//ff:what parseMeminfo — MemTotal and MemAvailable in bytes
package idle

import (
	"strconv"
	"strings"
)

func parseMeminfo(s string) (uint64, uint64) {
	var total, avail uint64
	for _, line := range strings.Split(s, "\n") {
		f := strings.Fields(line)
		if len(f) < 2 {
			continue
		}
		n, _ := strconv.ParseUint(f[1], 10, 64)
		switch f[0] {
		case "MemTotal:":
			total = n * 1024
		case "MemAvailable:":
			avail = n * 1024
		}
	}
	return total, avail
}
