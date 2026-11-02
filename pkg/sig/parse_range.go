//ff:func feature=sig type=decoder control=sequence
//ff:what parseRange — `LO..HI`, `..HI` or `LO..`
package sig

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func parseRange(s string) (int64, int64, error) {
	a, b, ok := strings.Cut(s, "..")
	if !ok || a == "" && b == "" {
		return 0, 0, fmt.Errorf("range %q", s)
	}
	lo, hi := int64(math.MinInt64), int64(math.MaxInt64)
	var err1, err2 error
	if a != "" {
		lo, err1 = strconv.ParseInt(a, 10, 64)
	}
	if b != "" {
		hi, err2 = strconv.ParseInt(b, 10, 64)
	}
	if err1 != nil || err2 != nil || lo >= hi {
		return 0, 0, fmt.Errorf("range %q", s)
	}
	return lo, hi, nil
}
