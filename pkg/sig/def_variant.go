//ff:func feature=sig type=builder control=sequence
//ff:what Def.variant — the plan for one code of the when field: that code alone there, plus the block's fields
package sig

import "slices"

func (d *Def) variant(e *Entity, base plan, on *Field, c uint64) plan {
	p := plan{fields: slices.Clone(base.fields), codes: slices.Clone(base.codes)}
	p.codes[slices.Index(p.fields, on)] = []uint64{c}
	if w := d.whenOf(on, c); w != nil {
		p.fields = append(p.fields, w.Fields...)
		p.codes = append(p.codes, e.fieldCodes(w.Fields)...)
	}
	return p
}
