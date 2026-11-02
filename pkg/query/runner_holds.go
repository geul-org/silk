//ff:func feature=query type=rule control=iteration dimension=1
//ff:what runner.holds — the silk conditions on a document's entities: yes for any entity, else unknown for any
package query

func (r *runner) holds(p *Plan, seq uint32) (bool, bool, error) {
	d, err := r.doc(seq)
	if err != nil {
		return false, false, err
	}
	unknown := false
	for _, e := range p.Def.Entities(r.ix.Grammar, d) {
		y, u, err := p.Def.Holds(e, p.Cond.Silk)
		if err != nil {
			return false, false, err
		}
		if y {
			return true, false, nil
		}
		unknown = unknown || u
	}
	return false, unknown, nil
}
