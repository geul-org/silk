//ff:func feature=index type=helper control=iteration dimension=1
//ff:what refs — add every file a manifest names to keep
package index

func refs(m Manifest, keep map[string]bool) {
	keep[manifestName(m.Generation)] = true
	keep[m.Live] = true
	keep[dirGrammar+"/"+m.Grammar+".json"] = true
	for _, d := range m.Docs {
		keep[d.File] = true
	}
	for _, s := range append([]Section{m.Paths}, m.Sections...) {
		keep[dirDecl+"/"+s.Decl+".json"] = true
		addSegs(keep, s.Runs)
		addSegs(keep, s.Appends)
	}
}
