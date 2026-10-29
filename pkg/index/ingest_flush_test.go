//ff:func feature=index type=writer control=sequence
//ff:what TestIngestFlush — every non-empty bucket written, packs closed; a write error is reported
package index

import "testing"

func TestIngestFlush(t *testing.T) {
	g, _, dir, docs := testIngest(t, 1<<20)
	g.doc(docs[0], 0, fill(7, true))
	if err := g.flush(); err != nil || len(g.out[0]) != 1 || g.files != nil {
		t.Fatal(err)
	}
	g2, _, _, _ := testIngest(t, 1<<20)
	g2.dir = dir
	g2.doc(docs[0], 0, fill(7, true))
	block(dir, dirSeg)
	if err := g2.flush(); err == nil {
		t.Fatal("blocked segment directory")
	}
}
