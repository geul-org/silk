//ff:func feature=query type=decoder control=iteration dimension=1
//ff:what ParseMatch — `LABEL:2.0=@ko,1=Q42` as a condition: a layout name, then field paths and tokens
package query

import (
	"fmt"
	"strings"
)

// ParseMatch reads `NAME:PATH=TOKEN[,PATH=TOKEN...]`. NAME is an Edge
// layout SILK knows (spec.LayoutOf); PATH is a field number or `field.0`.
func ParseMatch(s string) (Cond, error) {
	name, rest, ok := strings.Cut(s, ":")
	kind, codec, known := kindOf(name)
	if !ok || !known {
		return Cond{}, fmt.Errorf("%w: %q is not KIND:PATH=VALUE with a known kind", ErrSyntax, s)
	}
	c := Cond{Text: s, Name: name, Kind: kind, Codec: codec}
	for _, part := range strings.Split(rest, ",") {
		f, err := parseFieldCond(part)
		if err != nil {
			return Cond{}, err
		}
		c.Fields = appendField(c.Fields, f.Path, f.Value)
	}
	if len(c.Fields) == 0 {
		return Cond{}, fmt.Errorf("%w: %q fixes no field", ErrSyntax, s)
	}
	return c, nil
}
