//ff:func feature=index type=builder control=iteration dimension=1
//ff:what newIngest — an ingest for a spec: a bucket per tag carrying its TAG WORD, key widths, tag hash and minimum partition; pack and Issuer numbers from the manifest
package index

import (
	"os"

	"github.com/park-jun-woo/silk/pkg/seg"
	"github.com/park-jun-woo/silk/pkg/spec"
)

// newIngest writes segment files under dir and names them through m.
func newIngest(dir string, m *Manifest, sp spec.Spec, sorted []bool, batch uint64) *ingest {
	g := &ingest{dir: dir, m: m, sp: sp, table: sp.Table(), sorted: sorted, batch: max(batch, 1),
		out: make([][]SegRef, len(sp.Tags)), files: map[string]*os.File{}, packIx: map[string]uint32{}, issIx: map[string]uint32{}}
	for _, t := range sp.Tags {
		info := seg.Info{Tag: spec.TagWord(t.Kind, t.Codec), PartWidth: t.Partition.Width, SortWidth: t.Sort.Width, Spec: t.Hash()}
		g.buckets = append(g.buckets, seg.NewBucket(info, t.MinPartition))
	}
	for i, p := range m.Packs {
		g.packIx[p.Set+"/"+p.Hash] = uint32(i)
	}
	for i, s := range m.Issuers {
		g.issIx[s] = uint32(i)
	}
	return g
}
