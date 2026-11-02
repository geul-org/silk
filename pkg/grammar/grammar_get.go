//ff:func feature=grammar type=scanner control=sequence
//ff:what Grammar.Get — the field at path p of rec; false when the record has no such field
package grammar

func (g *Grammar) Get(rec []byte, p Path) (Val, bool) {
	return g.walk(nil, rec, p)
}
