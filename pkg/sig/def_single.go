//ff:func feature=sig type=builder control=sequence
//ff:what Def.Single — one row for an entity: a multi-valued field is `other`, so nothing spreads (the 32-bit half of a SIDX)
package sig

func (d *Def) Single(e *Entity) []byte {
	p := plan{fields: d.Common, codes: e.fieldCodes(d.Common)}
	if on := d.whenField(); on != nil {
		if c := e.codes(on); len(c) == 1 {
			p = d.variant(e, p, on, c[0])
		}
	}
	return p.limit(0).emit(nil, make([]byte, (d.Bits+7)/8), 0)[0]
}
