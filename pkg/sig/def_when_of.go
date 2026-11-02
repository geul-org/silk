//ff:func feature=sig type=helper control=iteration dimension=1
//ff:what Def.whenOf — the when block of a code
package sig

func (d *Def) whenOf(on *Field, c uint64) *When {
	for _, w := range d.Whens {
		if w.On == on && w.Entry.Code<<(on.Width-w.Entry.Len) == c {
			return w
		}
	}
	return nil
}
