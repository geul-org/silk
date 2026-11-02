//ff:func feature=grammar type=builder control=iteration dimension=1
//ff:what Doc.loadRefs — read every REFERENCE-NODE binding of the document
package grammar

func (d *Doc) loadRefs(g *Grammar) {
	d.refs = map[string][]byte{}
	for _, r := range d.Recs {
		if r.Tag != TagRefNode {
			continue
		}
		s, ok1 := g.Get(d.At(r), Path{1})
		o, ok2 := g.Get(d.At(r), Path{2})
		if ok1 && ok2 {
			d.refs[string(s.B)] = o.B
		}
	}
}
