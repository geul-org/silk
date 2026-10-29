//ff:func feature=query type=decoder control=sequence
//ff:what ParseTriple — `P106=Q169470` (either side may be *) as a Triple with canonical record bytes
package query

import (
	"fmt"
	"strings"
)

// ParseTriple reads `PRED=OBJ`; tokens are Qn, Pn, #n or * (any). At least
// one side must be given.
func ParseTriple(s string) (Triple, error) {
	p, o, ok := strings.Cut(s, "=")
	if !ok {
		return Triple{}, fmt.Errorf("%w: %q is not PRED=OBJ", ErrSyntax, s)
	}
	pred, err := encodeToken(strings.TrimSpace(p))
	if err != nil {
		return Triple{}, err
	}
	obj, err := encodeToken(strings.TrimSpace(o))
	if err != nil {
		return Triple{}, err
	}
	if pred == nil && obj == nil {
		return Triple{}, fmt.Errorf("%w: %q matches every TRIPLE", ErrSyntax, s)
	}
	return Triple{Text: s, Pred: pred, Obj: obj}, nil
}
