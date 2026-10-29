//ff:func feature=index type=helper control=iteration dimension=1
//ff:what findInputs — the section of a name and the references of the given files, false unless every file is still one of its segments
package index

// findInputs checks a checkpoint against the current sections.
func findInputs(secs []Section, name string, files []string) (int, []SegRef, bool) {
	for i, s := range secs {
		if s.Name != name {
			continue
		}
		refs := pickRefs(append(append([]SegRef{}, s.Runs...), s.Appends...), files)
		return i, refs, len(refs) == len(files) && len(files) > 0
	}
	return 0, nil, false
}
