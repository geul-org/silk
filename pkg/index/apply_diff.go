//ff:func feature=index type=builder control=iteration dimension=1
//ff:what applyDiff — index added versions (appended segments; all versions for rebuilt tags), update sections, kill removed versions, write the new generation
package index

import (
	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/spec"
)

// applyDiff gives added versions sequence numbers after the last one.
func applyDiff(dir string, m *Manifest, sp spec.Spec, a *archive.Archive, old []DocEntry, live Bitmap, d docDiff, rebuild []bool, o Options) (Report, error) {
	g := newIngest(dir, m, sp, rebuild, o.batch())
	if err := addKept(g, d.kept, rebuild); err != nil {
		return Report{}, err
	}
	keptRead := g.read
	first := uint32(len(old))
	entries, err := addDocs(g, d.added, first)
	if err != nil {
		return Report{}, err
	}
	if err := g.flush(); err != nil {
		return Report{}, err
	}
	m.Sections = nextSections(m.Sections, sp, rebuild, g.out)
	for _, seq := range d.killed {
		live.Put(seq, false)
		addDead(m.Sections, old[seq].Counts, rebuild)
	}
	for i := range entries {
		live.Put(first+uint32(i), true)
	}
	rep := Report{Docs: live.Count(), New: len(entries), Kept: len(d.kept), Killed: len(d.killed),
		NewBytes: g.read - keptRead, Read: g.read, Written: g.written}
	return rep, writeGeneration(dir, m, sp, a, entries, live, &rep)
}
