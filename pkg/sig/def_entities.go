//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what Def.Entities — group a document's TRIPLEs on the defined predicates by Subject, in first-seen order
package sig

import "github.com/geul-org/silk/pkg/grammar"

func (d *Def) Entities(g *grammar.Grammar, doc *grammar.Doc) []*Entity {
	var out []*Entity
	by := map[string]*Entity{}
	for _, r := range doc.Recs {
		if r.Tag != grammar.TagTriple {
			continue
		}
		s, p, o, ok := d.statement(g, doc.At(r))
		if !ok {
			continue
		}
		e := by[string(s)]
		if e == nil {
			e = &Entity{Subject: s, Objects: map[string][][]byte{}}
			by[string(s)] = e
			out = append(out, e)
		}
		e.Objects[string(p)] = append(e.Objects[string(p)], o)
	}
	return out
}
