//ff:func feature=sig type=decoder control=sequence
//ff:what gedcomYears — the last number of a GEDCOM date value, minus for BCE
package sig

import (
	"strconv"
	"strings"
)

func gedcomYears(s string) (int64, int64, bool) {
	f := strings.Fields(s)
	neg := len(f) > 0 && f[len(f)-1] == "BCE"
	if neg {
		f = f[:len(f)-1]
	}
	if len(f) == 0 {
		return 0, 0, false
	}
	n, err := strconv.ParseInt(f[len(f)-1], 10, 64)
	if neg {
		n = -n
	}
	return n, n, err == nil
}
