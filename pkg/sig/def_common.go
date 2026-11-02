//ff:func feature=sig type=helper control=iteration dimension=1
//ff:what Def.common — a common field by name
package sig

func (d *Def) common(name string) *Field {
	for _, f := range d.Common {
		if f.Name == name {
			return f
		}
	}
	return nil
}
