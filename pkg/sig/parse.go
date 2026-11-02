//ff:func feature=sig type=decoder control=iteration dimension=1
//ff:what Parse — read a definition: `silk NAME BITS`, `field NAME OFFSET WIDTH TRIPLE PREDICATE KIND`, `code …` under a field, `when FIELD=CODE`; everything explicit, nothing assigned
package sig

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func Parse(text string) (*Def, error) {
	p := &parser{d: &Def{Preds: map[string]bool{}, Hash: sha256.Sum256([]byte(text))}}
	for n, line := range strings.Split(text, "\n") {
		if err := p.line(strings.Fields(stripComment(line))); err != nil {
			return nil, fmt.Errorf("%w: line %d: %v", ErrDef, n+1, err)
		}
	}
	if p.d.Bits == 0 {
		return nil, fmt.Errorf("%w: no silk line", ErrDef)
	}
	if err := p.d.check(); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDef, err)
	}
	return p.d, nil
}
