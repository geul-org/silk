//ff:func feature=index type=builder control=iteration dimension=1
//ff:what nextSections — the sections of a spec after an ingest: a rebuilt tag gets its new runs; a kept tag keeps its section and gains the new appended segments
package index

import (
	"encoding/hex"

	"github.com/geul-org/silk/pkg/spec"
)

// nextSections orders sections as sp.Tags. Sections of tags no longer in
// the spec are dropped.
func nextSections(old []Section, sp spec.Spec, rebuild []bool, out [][]SegRef) []Section {
	res := make([]Section, 0, len(sp.Tags))
	for i, t := range sp.Tags {
		h := t.Hash()
		s := Section{Name: t.Name, Tag: spec.TagWord(t.Kind, t.Codec), Hash: hex.EncodeToString(h[:]), Runs: []SegRef{}, Appends: []SegRef{}}
		if rebuild[i] {
			s.Runs = append(s.Runs, out[i]...)
		} else {
			s = sectionNamed(old, t.Name)
			s.Appends = append(s.Appends, out[i]...)
		}
		res = append(res, s)
	}
	return res
}
