//ff:func feature=archive type=helper control=iteration dimension=1
//ff:what Archive.Set — a set by name; nil when absent
package archive

func (a *Archive) Set(name string) *Set {
	for _, s := range a.Sets {
		if s.Name == name {
			return s
		}
	}
	return nil
}
