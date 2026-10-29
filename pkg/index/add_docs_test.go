//ff:func feature=index type=builder control=sequence
//ff:what TestAddDocs — entries for every document with sequence numbers from first; a failing document stops
package index

import "testing"

func TestAddDocs(t *testing.T) {
	g, _, _, docs := testIngest(t, 1<<20)
	es, err := addDocs(g, docs, 10)
	if err != nil || len(es) != 4 || len(es[3].Counts) == 0 {
		t.Fatal(err)
	}
	bad := docs[0]
	bad.Path = "/nonexistent"
	if _, err := addDocs(g, []liveDoc{bad}, 0); err == nil {
		t.Fatal("unreadable document")
	}
}
