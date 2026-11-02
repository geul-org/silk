//ff:func feature=cli type=decoder control=iteration dimension=1
//ff:what parseWidths — comma separated numbers; empty is none
package cli

import (
	"fmt"
	"strconv"
	"strings"
)

func parseWidths(s string) ([]int, error) {
	if s == "" {
		return nil, nil
	}
	var out []int
	for _, p := range strings.Split(s, ",") {
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("width %q", p)
		}
		out = append(out, n)
	}
	return out, nil
}
