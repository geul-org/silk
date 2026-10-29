//ff:func feature=index type=builder control=sequence
//ff:what TestIngestDoc — rows for masked tags only and record counts for every kind; unreadable or malformed documents and a failing flush fail
package index

import "testing"

func TestIngestDoc(t *testing.T) {
	g, _, dir, docs := testIngest(t, 1<<20)
	mask := fill(7, false)
	mask[0] = true
	counts, err := g.doc(docs[0], 5, mask)
	if err != nil || g.buckets[0].Rows == 0 || g.buckets[2].Rows != 0 || countOf(counts, 0x1004) == 0 || uint64(countOf(counts, 0x100A)) != g.buckets[0].Rows {
		t.Fatal(err, counts)
	}
	bad := docs[0]
	bad.Path = "/nonexistent"
	if _, err := g.doc(bad, 0, mask); err == nil {
		t.Fatal("unreadable")
	}
	bad = docs[0]
	bad.Doc.Off += 2
	if _, err := g.doc(bad, 0, mask); err == nil {
		t.Fatal("malformed")
	}
	g.batch = 1
	block(dir, dirSeg)
	if _, err := g.doc(docs[1], 6, mask); err == nil {
		t.Fatal("flush")
	}
}
