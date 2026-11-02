//ff:func feature=sig type=decoder control=sequence
//ff:what Def.statement — a TRIPLE's Subject, Predicate and Object when its predicate is defined and its rank is not deprecated
package sig

import (
	"encoding/binary"

	"github.com/geul-org/silk/pkg/grammar"
)

func (d *Def) statement(g *grammar.Grammar, rec []byte) ([]byte, []byte, []byte, bool) {
	p, ok := g.Get(rec, grammar.Path{1})
	if !ok || !d.Preds[string(p.B)] {
		return nil, nil, nil, false
	}
	s, ok1 := g.Get(rec, grammar.Path{2})
	o, ok2 := g.Get(rec, grammar.Path{3})
	r, ok3 := g.Get(rec, grammar.Path{4})
	if !ok1 || !ok2 || !ok3 || binary.BigEndian.Uint16(r.B) == rankDeprecated {
		return nil, nil, nil, false
	}
	return s.B, p.B, o.B, true
}
