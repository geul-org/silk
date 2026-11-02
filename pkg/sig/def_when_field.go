//ff:func feature=sig type=helper control=sequence
//ff:what Def.whenField — the common field the when blocks hang on; nil without blocks
package sig

func (d *Def) whenField() *Field {
	if len(d.Whens) == 0 {
		return nil
	}
	return d.Whens[0].On
}
