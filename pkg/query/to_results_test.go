//ff:func feature=query type=builder control=sequence
//ff:what TestToResults — results named through the document table, sorted by Issuer, document and version
package query

import (
	"testing"

	"github.com/geul-org/silk/pkg/index"
)

func TestToResults(t *testing.T) {
	ix := &index.Index{M: index.Manifest{Issuers: []string{"0x0002", "0x0001"}},
		Docs: []index.DocEntry{{Issuer: 0, ID: 1, Version: 1}, {Issuer: 1, ID: 2, Version: 2}, {Issuer: 1, ID: 2, Version: 1}, {Issuer: 1, ID: 1, Version: 1}}}
	rs := toResults(ix, []docHit{{0, []uint64{1}}, {1, nil}, {2, nil}, {3, nil}})
	if rs[0].Doc != 1 || rs[1].Version != 1 || rs[2].Version != 2 || rs[3].Issuer != "0x0002" || rs[3].Edges[0] != 1 {
		t.Fatal(rs)
	}
}
