//ff:func feature=index type=builder control=sequence
//ff:what TestAppendDocRows — full rows (partition slot, sort slot, sequence, offset) for every indexed record of a document
package index

import (
	"bytes"
	"context"
	"testing"

	"github.com/park-jun-woo/silk/pkg/seg"
)

func TestAppendDocRows(t *testing.T) {
	a, _, _ := built(t)
	docs, _ := liveDocs(context.Background(), a)
	g := newIngest("d", &Manifest{}, smallSpec(), fill(7, false), 1)
	b, _ := g.readDoc(docs[0])
	want := appendDocRows(make([][]byte, 7), g, b, 9)
	if len(want[0]) == 0 || len(want[0])%20 != 0 {
		t.Fatal(len(want[0]))
	}
	doc, off := seg.SeqOf(want[0][:20])
	rec := b[off:]
	if doc != 9 || !bytes.Equal(want[0][:12], g.sp.Tags[0].Extract(nil, rec[:recLenAt(b, int(off))])) {
		t.Fatal(doc, off)
	}
}
