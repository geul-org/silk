//ff:func feature=index type=builder control=iteration dimension=1
//ff:what planNamed — the due tasks of a manifest with output file names reserved and each tag's minimum partition
package index

import "github.com/geul-org/silk/pkg/spec"

// planNamed completes Plan's tasks for execution.
func planNamed(m *Manifest, sp spec.Spec, o Options) []Task {
	ts := Plan(*m, o.Policy, o.Full)
	for i := range ts {
		ts[i].Output = m.newName(dirSeg, ".sseg")
		t, _ := sp.Find(m.Sections[ts[i].Section].Tag)
		ts[i].Min = t.MinPartition
	}
	return ts
}
