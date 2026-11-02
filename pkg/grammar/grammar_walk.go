//ff:func feature=grammar type=scanner control=iteration dimension=1
//ff:what Grammar.walk — follow a path field by field
package grammar

func (g *Grammar) walk(d *Doc, rec []byte, p Path) (Val, bool) {
	v := Val{B: rec}
	for k, i := range p {
		if v.Raw {
			return Val{}, false
		}
		if k > 0 && d != nil {
			v.B = d.Resolve(g, v.B)
		}
		var ok bool
		if v, ok = g.field(v.B, i); !ok {
			return Val{}, false
		}
	}
	return v, true
}
