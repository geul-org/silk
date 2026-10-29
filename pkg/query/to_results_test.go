//ff:func feature=query type=builder control=sequence
//ff:what TestToResults — sorted by Issuer, then document, then version
package query

import "testing"

func TestToResults(t *testing.T) {
	rs := toResults(map[docKey][]uint64{
		{issuer: "0x0002", doc: 1, version: 1}: {1},
		{issuer: "0x0001", doc: 5, version: 2}: {2},
		{issuer: "0x0001", doc: 5, version: 1}: {3},
		{issuer: "0x0001", doc: 3, version: 9}: {4},
	})
	if len(rs) != 4 || rs[0].Edges[0] != 4 || rs[1].Edges[0] != 3 || rs[2].Edges[0] != 2 || rs[3].Issuer != "0x0002" {
		t.Fatal(rs)
	}
}
