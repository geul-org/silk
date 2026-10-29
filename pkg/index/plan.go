//ff:func feature=index type=rule control=iteration dimension=1
//ff:what Plan — the merges a manifest's sections are due under a policy (at most one per section; full asks for one sorted run each)
package index

// Plan returns tasks without output names.
func Plan(m Manifest, p Policy, full bool) []Task {
	var out []Task
	for i, s := range m.Sections {
		if t, ok := planSection(i, s, p.resolve(), full); ok {
			out = append(out, t)
		}
	}
	return out
}
