//ff:func feature=index type=builder control=iteration dimension=1
//ff:what resetManifest — the manifest of a build from scratch: declarations kept, everything built forgotten
package index

import "slices"

func resetManifest(m Manifest) Manifest {
	out := Manifest{Version: m.Version, Generation: m.Generation, Grammar: m.Grammar, Paths: Section{Name: "paths"}, Sections: slices.Clone(m.Sections), NextFile: m.NextFile}
	for i := range out.Sections {
		out.Sections[i].Pending, out.Sections[i].Runs, out.Sections[i].Appends = true, nil, nil
	}
	return out
}
