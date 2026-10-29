//ff:func feature=query type=builder control=sequence
//ff:what TestGroupHits — matched candidates grouped per document, Edge Ids in order, unmatched ones skipped
package query

import "testing"

func TestGroupHits(t *testing.T) {
	cs := []candidate{{doc: 1}, {doc: 1}, {doc: 2}, {doc: 3}}
	hs := groupHits(cs, []uint64{7, 8, 9, 10}, []bool{true, true, false, true})
	if len(hs) != 2 || len(hs[0].edges) != 2 || hs[0].edges[1] != 8 || hs[1].doc != 3 {
		t.Fatal(hs)
	}
}
