//ff:func feature=sig type=decoder control=sequence
//ff:what bucketCode — one number for one half-open range LO..HI; either end may be open
package sig

import "errors"

func bucketCode(name, code string, rest []string, width int) (*Entry, error) {
	n, err := parseNumber(code)
	if err != nil || len(rest) != 1 {
		return nil, errors.New("code NAME NUMBER LO..HI")
	}
	lo, hi, err := parseRange(rest[0])
	return &Entry{Name: name, Code: n, Len: width, Lo: lo, Hi: hi}, err
}
