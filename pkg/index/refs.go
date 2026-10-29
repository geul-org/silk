//ff:func feature=index type=helper control=iteration dimension=1
//ff:what refs — add every file a manifest names (itself, spec, document tables, live bitmap, segments) to a set
package index

// refs marks m's files in keep (relative, slash-separated).
func refs(m Manifest, keep map[string]bool) {
	keep[manifestName(m.Generation)] = true
	keep[specName(m.Spec)] = true
	keep[m.Live] = true
	for _, d := range m.Docs {
		keep[d.File] = true
	}
	for _, s := range m.Sections {
		for _, r := range append(append([]SegRef{}, s.Runs...), s.Appends...) {
			keep[r.File] = true
		}
	}
}
