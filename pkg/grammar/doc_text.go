//ff:func feature=grammar type=decoder control=selection
//ff:what Doc.Text — a text Edge's language record and text bytes: Object is at field 2, a LOCAL resolved; false for anything else
package grammar

func (d *Doc) Text(g *Grammar, rec []byte) ([]byte, []byte, bool) {
	o, ok := g.Get(rec, Path{2})
	if !ok || o.Raw {
		return nil, nil, false
	}
	b := d.Resolve(g, o.B)
	switch Target(b) {
	case TagText, TagTextOdd:
		v, ok := g.Get(b, Path{0})
		return nil, v.B, ok
	case TagLangText, TagLangOdd:
		l, ok1 := g.Get(b, Path{0})
		v, ok2 := g.Get(b, Path{1})
		return l.B, v.B, ok1 && ok2
	}
	return nil, nil, false
}
