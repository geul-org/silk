//ff:func feature=query type=decoder control=iteration dimension=1
//ff:what parseFieldCond — `2.0=@ko` as a field path and record bytes (nil for *)
package query

import (
	"fmt"
	"strconv"
	"strings"
)

// parseFieldCond reads one PATH=TOKEN.
func parseFieldCond(s string) (FieldCond, error) {
	p, tok, ok := strings.Cut(strings.TrimSpace(s), "=")
	if !ok {
		return FieldCond{}, fmt.Errorf("%w: %q is not PATH=VALUE", ErrSyntax, s)
	}
	var path []int
	for _, x := range strings.Split(p, ".") {
		n, err := strconv.Atoi(x)
		if err != nil || n < 0 {
			return FieldCond{}, fmt.Errorf("%w: field path %q", ErrSyntax, p)
		}
		path = append(path, n)
	}
	v, err := encodeToken(strings.TrimSpace(tok))
	return FieldCond{Path: path, Value: v}, err
}
