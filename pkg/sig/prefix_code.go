//ff:func feature=sig type=decoder control=sequence
//ff:what prefixCode — leading bits written out; without values it only names a group for queries
package sig

import (
	"errors"
	"strconv"
	"strings"
)

func prefixCode(name, code string, rest []string) (*Entry, error) {
	bits, ok := strings.CutPrefix(code, "0b")
	if !ok || bits == "" || len(bits) > 64 || strings.Trim(bits, "01") != "" {
		return nil, errors.New("code NAME 0bBITS VALUE…")
	}
	n, _ := strconv.ParseUint(bits, 2, 64)
	vals, err := parseValues(rest)
	return &Entry{Name: name, Code: n, Len: len(bits), Values: vals}, err
}
