//ff:func feature=query type=decoder control=sequence
//ff:what ParseTriple — `P106=Q169470` (either side may be *) as a TRIPLE condition on Predicate (field 1) and Object (field 3)
package query

import (
	"fmt"
	"strings"
)

// ParseTriple reads `PRED=OBJ`; at least one side must be given.
func ParseTriple(s string) (Cond, error) {
	p, o, ok := strings.Cut(s, "=")
	if !ok {
		return Cond{}, fmt.Errorf("%w: %q is not PRED=OBJ", ErrSyntax, s)
	}
	pred, err := encodeToken(strings.TrimSpace(p))
	if err != nil {
		return Cond{}, err
	}
	obj, err := encodeToken(strings.TrimSpace(o))
	if err != nil {
		return Cond{}, err
	}
	if pred == nil && obj == nil {
		return Cond{}, fmt.Errorf("%w: %q matches every TRIPLE", ErrSyntax, s)
	}
	c := Cond{Text: s, Name: "TRIPLE", Kind: 2, Codec: 0x0A}
	c.Fields = appendField(appendField(nil, []int{1}, pred), []int{3}, obj)
	return c, nil
}
