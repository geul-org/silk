//ff:func feature=index type=builder control=iteration dimension=1
//ff:what ingest.doc — index one document: validate it, count its records per TAG kind, add a row for each record of a tag in mask; flush full buckets
package index

import (
	"encoding/binary"

	"github.com/park-jun-woo/silk/pkg/seg"
)

// doc gives document d the sequence number seq. mask selects the tags that
// get rows (all for a new version, rebuilt tags only for a kept one).
func (g *ingest) doc(d liveDoc, seq uint32, mask []bool) ([]TagCount, error) {
	b, err := g.readDoc(d)
	if err != nil {
		return nil, err
	}
	if err := checkDoc(b); err != nil {
		return nil, err
	}
	var counts []TagCount
	for p := 0; p < len(b); {
		n := recLenAt(b, p)
		rec := b[p : p+n]
		tw := binary.BigEndian.Uint16(rec) & 0x187F
		counts = addCount(counts, tw)
		if i := g.table[tw]; i >= 0 && mask[i] {
			t := g.sp.Tags[i]
			g.key = t.Extract(g.key[:0], rec)
			g.buckets[i].Add(g.key[:t.Partition.Width], seg.AppendSeq(g.key[t.Partition.Width:], seq, uint32(p)))
		}
		p += n
	}
	return counts, g.flushFull()
}
