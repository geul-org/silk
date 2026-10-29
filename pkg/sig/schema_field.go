//ff:func feature=sig type=helper control=iteration dimension=1
//ff:what Schema.field — a field by name
package sig

// field looks name up in s.Fields.
func (s Schema) field(name string) (Field, bool) {
	for _, f := range s.Fields {
		if f.Name == name {
			return f, true
		}
	}
	return Field{}, false
}
