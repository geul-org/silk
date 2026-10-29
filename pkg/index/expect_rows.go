//ff:func feature=index type=builder control=iteration dimension=1
//ff:what expectRows — the rows each section should hold, rebuilt from the live documents (partition slot, sort slot, sequence, offset)
package index

import (
	"github.com/geul-org/silk/pkg/seg"
)

// expectRows reads every document with an ingest that never writes.
func expectRows(ix *Index, docs []liveDoc, seqs []uint32) ([][]byte, error) {
	g := newIngest(ix.Dir, &Manifest{}, ix.Spec, fill(len(ix.Spec.Tags), false), ^uint64(0))
	want := make([][]byte, len(ix.Spec.Tags))
	for i, d := range docs {
		b, err := g.readDoc(d)
		if err != nil {
			return nil, err
		}
		if err := checkDoc(b); err != nil {
			return nil, err
		}
		want = appendDocRows(want, g, b, seqs[i])
	}
	for i := range want {
		t := ix.Spec.Tags[i]
		want[i] = seg.SortRows(want[i], t.Partition.Width+t.Sort.Width+8)
	}
	return want, g.flush()
}
