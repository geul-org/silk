//ff:func feature=index type=reader control=iteration dimension=1
//ff:what Index.loadSections — open each built section and read its declaration
package index

func (ix *Index) loadSections() error {
	for _, s := range ix.Manifest.Sections {
		d, err := readDecl(ix.Dir, s.Decl)
		if err != nil {
			return err
		}
		ix.Decls[s.Name] = d
		if s.Pending {
			continue
		}
		fs, err := openSegs(segPaths(ix.Dir, s))
		if err != nil {
			return err
		}
		ix.secs[s.Name] = append(ix.secs[s.Name], &Sec{Section: s, Files: fs})
	}
	return nil
}
