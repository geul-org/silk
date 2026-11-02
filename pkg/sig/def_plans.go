//ff:func feature=sig type=builder control=iteration dimension=1
//ff:what Entity.plans — one plan per row variant: each code of a when field picks its block's fields
package sig

func (d *Def) plans(e *Entity) []plan {
	base := plan{fields: d.Common, codes: e.fieldCodes(d.Common)}
	on := d.whenField()
	if on == nil {
		return []plan{base}
	}
	var out []plan
	for _, c := range e.codes(on) {
		out = append(out, d.variant(e, base, on, c))
	}
	return out
}
