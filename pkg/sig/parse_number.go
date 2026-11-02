//ff:func feature=sig type=decoder control=sequence
//ff:what parseNumber — a code number in decimal or 0x hex
package sig

import (
	"strconv"
	"strings"
)

func parseNumber(s string) (uint64, error) {
	if h, ok := strings.CutPrefix(s, "0x"); ok {
		return strconv.ParseUint(h, 16, 64)
	}
	return strconv.ParseUint(s, 10, 64)
}
