//ff:func feature=index type=writer control=sequence
//ff:what TestIngestFlushOne — an empty bucket writes nothing; a full one becomes a named segment with its rows and bytes; a write error is reported
package index

import (
	"testing"
)

func TestIngestFlushOne(t *testing.T) {
	g, m, dir, docs := testIngest(t, 1<<20)
	if err := g.flushOne(0); err != nil || len(g.out[0]) != 0 {
		t.Fatal(err)
	}
	g.doc(docs[0], 0, fill(7, true))
	rows := g.buckets[0].Rows
	if err := g.flushOne(0); err != nil || g.out[0][0].Rows != rows || g.written != g.out[0][0].Bytes || m.NextFile != 1 {
		t.Fatal(err, g.out[0])
	}
	g.doc(docs[1], 1, fill(7, true))
	block(dir, dirSeg)
	if err := g.flushOne(0); err == nil {
		t.Fatal("blocked segment directory")
	}
}
