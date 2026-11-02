//ff:func feature=sig type=decoder control=iteration dimension=1
//ff:what When.field — a block's field by name
package sig

func (w *When) field(name string) *Field {
	for _, f := range w.Fields {
		if f.Name == name {
			return f
		}
	}
	return nil
}
