//ff:func feature=grammar type=builder control=sequence
//ff:what Doc.Resolve — the Object a REFERENCE-NODE of this document binds to a LOCAL record; any other record as is
package grammar

func (d *Doc) Resolve(g *Grammar, rec []byte) []byte {
	if len(rec) < 2 || Target(rec) != TagLocal {
		return rec
	}
	if d.refs == nil {
		d.loadRefs(g)
	}
	if o, ok := d.refs[string(rec)]; ok {
		return o
	}
	return rec
}
