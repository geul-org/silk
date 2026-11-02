//ff:func feature=cli type=decoder control=selection
//ff:what parseSize — bytes from `123`, `64M`, `2G`
package cli

import (
	"fmt"
	"strconv"
	"strings"
)

func parseSize(s string) (int64, error) {
	mult := int64(1)
	switch {
	case strings.HasSuffix(s, "K"):
		mult = 1 << 10
	case strings.HasSuffix(s, "M"):
		mult = 1 << 20
	case strings.HasSuffix(s, "G"):
		mult = 1 << 30
	}
	n, err := strconv.ParseInt(strings.TrimRight(s, "KMG"), 10, 64)
	if err != nil || n < 0 {
		return 0, fmt.Errorf("size %q", s)
	}
	return n * mult, nil
}
