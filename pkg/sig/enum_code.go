//ff:func feature=sig type=decoder control=sequence
//ff:what enumCode — one number standing for one or more values
package sig

import "errors"

func enumCode(name, code string, rest []string, width int) (*Entry, error) {
	n, err := parseNumber(code)
	if err != nil || len(rest) == 0 {
		return nil, errors.New("code NAME NUMBER VALUE…")
	}
	vals, err := parseValues(rest)
	return &Entry{Name: name, Code: n, Len: width, Values: vals}, err
}
