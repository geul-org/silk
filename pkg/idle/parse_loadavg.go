//ff:func feature=idle type=decoder control=sequence
//ff:what parseLoadavg — the 1-minute load average, the first field of /proc/loadavg
package idle

import (
	"fmt"
	"strconv"
	"strings"
)

// parseLoadavg reads "0.17 0.10 0.09 1/631 3321009".
func parseLoadavg(b []byte) (float64, error) {
	f := strings.Fields(string(b))
	if len(f) == 0 {
		return 0, fmt.Errorf("%w: empty loadavg", ErrUnsupported)
	}
	return strconv.ParseFloat(f[0], 64)
}
