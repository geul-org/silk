//ff:func feature=sig type=decoder control=iteration dimension=1
//ff:what Def.field — a field by name, with its when block; nil when there is none
package sig

func (d *Def) field(name string) (*Field, *When) {
	if f := d.common(name); f != nil {
		return f, nil
	}
	for _, w := range d.Whens {
		if f := w.field(name); f != nil {
			return f, w
		}
	}
	return nil, nil
}
