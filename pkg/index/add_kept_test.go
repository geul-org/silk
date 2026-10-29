//ff:func feature=index type=builder control=sequence
//ff:what TestAddKept — nothing is read without rebuilt tags; with one, kept documents get rows for it; a failing flush fails
package index

import (
	"testing"
)

func TestAddKept(t *testing.T) {
	g, _, dir, docs := testIngest(t, 1<<20)
	kept := []keptDoc{{seq: 0, doc: docs[0]}, {seq: 1, doc: docs[1]}}
	if err := addKept(g, kept, fill(7, false)); err != nil || g.read != 0 {
		t.Fatal(err)
	}
	mask := fill(7, false)
	mask[2] = true
	if err := addKept(g, kept, mask); err != nil || g.buckets[2].Rows == 0 || g.buckets[0].Rows != 0 {
		t.Fatal(err)
	}
	g.batch = 1
	block(dir, dirSeg)
	if err := addKept(g, kept, mask); err == nil {
		t.Fatal("blocked segment directory")
	}
}
