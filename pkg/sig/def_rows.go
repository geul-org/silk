//ff:func feature=sig type=builder control=iteration dimension=1
//ff:what Def.Rows — the rows of an entity: every code combination, limited when they pass MaxExpand; sorted and distinct
package sig

func (d *Def) Rows(e *Entity) [][]byte {
	ps := d.plans(e)
	n := 0
	for _, p := range ps {
		n += p.product()
	}
	var out [][]byte
	for _, p := range ps {
		out = p.bounded(n, len(ps)).emit(out, make([]byte, (d.Bits+7)/8), 0)
	}
	return sortRows(out)
}
