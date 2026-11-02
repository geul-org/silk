//ff:func feature=sig type=helper control=iteration dimension=1
//ff:what Field.entry — a code by name
package sig

func (f *Field) entry(name string) *Entry {
	for _, e := range f.Codes {
		if e.Name == name {
			return e
		}
	}
	return nil
}
