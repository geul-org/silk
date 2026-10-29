//ff:func feature=index type=writer control=sequence
//ff:what TestIngestFlushFull — only buckets at the batch size are written; a write error is reported
package index

import (
	"testing"
)

func TestIngestFlushFull(t *testing.T) {
	g, _, dir, docs := testIngest(t, 1<<20)
	g.doc(docs[0], 0, fill(7, true))
	g.batch = g.buckets[0].Rows
	if err := g.flushFull(); err != nil || len(g.out[0]) != 1 || g.buckets[0].Rows != 0 || len(g.out[6]) != 0 {
		t.Fatal(err, g.out)
	}
	g.batch = 1
	block(dir, dirSeg)
	if err := g.flushFull(); err == nil {
		t.Fatal("blocked segment directory")
	}
}
