//ff:func feature=index type=builder control=iteration dimension=1
//ff:what appendDocRows — the full rows (partition slot first) of one checked document, appended per tag
package index

import (
	"encoding/binary"

	"github.com/park-jun-woo/silk/pkg/seg"
)

// appendDocRows frames b's records like ingest.doc.
func appendDocRows(want [][]byte, g *ingest, b []byte, seq uint32) [][]byte {
	for p := 0; p < len(b); {
		n := recLenAt(b, p)
		if i := g.table[binary.BigEndian.Uint16(b[p:])&0x187F]; i >= 0 {
			want[i] = seg.AppendSeq(g.sp.Tags[i].Extract(want[i], b[p:p+n]), seq, uint32(p))
		}
		p += n
	}
	return want
}
