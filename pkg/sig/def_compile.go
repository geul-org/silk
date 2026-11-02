//ff:func feature=sig type=builder control=iteration dimension=1
//ff:what Def.Compile — the patterns of a condition list. A field under `when` needs that code; rows whose when field is somevalue or other carry no such field and stay candidates
package sig

import (
	"fmt"
	"slices"
)

func (d *Def) Compile(cs []Cond) ([]Pattern, error) {
	var common []Term
	var block []Term
	var when *When
	for _, c := range cs {
		f, w := d.field(c.Field)
		if f == nil {
			return nil, fmt.Errorf("%w: no field %q", ErrCond, c.Field)
		}
		t, err := f.term(c.Value)
		if err != nil {
			return nil, err
		}
		if w == nil {
			common = append(common, t)
			continue
		}
		if when != nil && when != w {
			return []Pattern{d.pattern(append(common, d.whenTerm(w.On, false)))}, nil
		}
		when, block = w, append(block, t)
	}
	if when == nil {
		return []Pattern{d.pattern(common)}, nil
	}
	a := d.pattern(append(append(slices.Clone(common), block...), d.exactTerm(when)))
	return []Pattern{a, d.pattern(append(common, d.whenTerm(when.On, false)))}, nil
}
