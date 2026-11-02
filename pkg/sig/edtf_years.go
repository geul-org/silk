//ff:func feature=sig type=decoder control=sequence
//ff:what edtfYears — an EDTF date's year span: `1952-03-11`, `1952`, `195X`, `19XX`, `-0043`
package sig

import (
	"strconv"
	"strings"
)

func edtfYears(s string) (int64, int64, bool) {
	neg := strings.HasPrefix(s, "-")
	y := strings.TrimPrefix(s, "-")
	if i := strings.IndexByte(y, '-'); i >= 0 {
		y = y[:i]
	}
	lo, err1 := strconv.ParseInt(strings.ReplaceAll(y, "X", "0"), 10, 64)
	hi, err2 := strconv.ParseInt(strings.ReplaceAll(y, "X", "9"), 10, 64)
	if err1 != nil || err2 != nil {
		return 0, 0, false
	}
	if neg {
		return -hi, -lo, true
	}
	return lo, hi, true
}
