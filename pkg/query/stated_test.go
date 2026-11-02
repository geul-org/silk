//ff:func feature=query type=rule control=iteration dimension=1
//ff:what stated — a document has a non-deprecated TRIPLE pred=obj, read with no silk code
package query

import (
	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/grammar"
)

func stated(g *grammar.Grammar, d fixture.Doc, pred, obj string) bool {
	for _, r := range d.Body {
		if r.Codec != 0x0A {
			continue
		}
		b := grammar.RecordOf(r)
		p, _ := g.Get(b, grammar.Path{1})
		o, _ := g.Get(b, grammar.Path{3})
		k, _ := g.Get(b, grammar.Path{4})
		if string(p.B) == string(fixture.Tok(pred)) && string(o.B) == string(fixture.Tok(obj)) && k.B[1] != 2 {
			return true
		}
	}
	return false
}
