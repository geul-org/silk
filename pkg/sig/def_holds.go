//ff:func feature=sig type=rule control=iteration dimension=1
//ff:what Def.Holds — the conditions read on an entity's statements: yes when each field has the asked code or the asked value itself; unknown when a miss could still be somevalue
package sig

import "fmt"

func (d *Def) Holds(e *Entity, cs []Cond) (bool, bool, error) {
	yes, unknown := true, false
	for _, c := range cs {
		f, w := d.field(c.Field)
		if f == nil {
			return false, false, fmt.Errorf("%w: no field %q", ErrCond, c.Field)
		}
		t, err := f.term(c.Value)
		if err != nil {
			return false, false, err
		}
		has, some := scoped(e, w, f, t, c.Value)
		yes, unknown = yes && has, unknown || !has && some
		if !has && !some {
			return false, false, nil
		}
	}
	return yes, unknown && !yes, nil
}
