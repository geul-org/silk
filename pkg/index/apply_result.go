//ff:func feature=index type=builder control=sequence
//ff:what applyResult — put a finished merge into the manifest: inputs leave the section, the output joins its runs (unless empty), dropped rows leave the dead count
package index

// applyResult edits m in memory; commit makes it a generation.
func applyResult(m *Manifest, r taskResult) {
	s := &m.Sections[r.Task.Section]
	gone := fileSet(r.Task.Inputs)
	s.Runs, s.Appends = without(s.Runs, gone), without(s.Appends, gone)
	if r.Out.Rows > 0 {
		s.Runs = append(s.Runs, r.Out)
	}
	s.Dead -= min(s.Dead, r.Dropped)
}
