//ff:func feature=archive type=decoder control=sequence
//ff:what parseHexWidth — `0x` and exactly n lowercase hex digits as a number
package archive

import (
	"strconv"
	"strings"
)

// parseHexWidth reads s as `0x` plus n lowercase hex digits.
func parseHexWidth(s string, n int) (uint64, bool) {
	h, ok := strings.CutPrefix(s, "0x")
	if !ok || len(h) != n || strings.ToLower(h) != h {
		return 0, false
	}
	v, err := strconv.ParseUint(h, 16, 64)
	return v, err == nil
}
