//ff:func feature=grammar type=scanner control=sequence
//ff:what Doc.Get — like Grammar.Get, resolving a LOCAL Reference through this document's REFERENCE-NODE when the path goes on inside it
package grammar

func (d *Doc) Get(g *Grammar, rec []byte, p Path) (Val, bool) {
	return g.walk(d, rec, p)
}
