//ff:func feature=index type=writer control=iteration dimension=1
//ff:what cloneDecls — a new index directory with the same declarations, nothing built
package index

func cloneDecls(from, to string, m Manifest) error {
	g, err := loadGrammar(from, m.Grammar)
	if err != nil {
		return err
	}
	seen := map[string]bool{}
	for _, s := range m.Sections {
		if seen[s.Decl] {
			continue
		}
		seen[s.Decl] = true
		d, err := readDecl(from, s.Decl)
		if err != nil {
			return err
		}
		if err := Declare(to, d, g); err != nil {
			return err
		}
	}
	return nil
}
