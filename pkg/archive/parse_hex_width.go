//ff:func feature=archive type=decoder control=sequence
//ff:what parseHexWidth — `0x` then exactly n lower-case hex digits
package archive

import (
	"strconv"
	"strings"
)

func parseHexWidth(s string, n int) (uint64, bool) {
	if len(s) != n+2 || !strings.HasPrefix(s, "0x") || strings.ToLower(s) != s {
		return 0, false
	}
	v, err := strconv.ParseUint(s[2:], 16, 64)
	return v, err == nil
}
