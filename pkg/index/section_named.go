//ff:func feature=index type=helper control=iteration dimension=1
//ff:what sectionNamed — the section of a name (zero Section when absent)
package index

// sectionNamed finds a section by tag name.
func sectionNamed(secs []Section, name string) Section {
	for _, s := range secs {
		if s.Name == name {
			return s
		}
	}
	return Section{}
}
