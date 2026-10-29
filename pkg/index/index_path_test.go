//ff:func feature=index type=helper control=sequence
//ff:what TestIndexPath — Issuer, 16-digit document ID and 8-digit version
package index

import "testing"

func TestIndexPath(t *testing.T) {
	ix := &Index{M: Manifest{Issuers: []string{"0x0001"}}, Docs: []DocEntry{{ID: 0x2a, Version: 3}}}
	if p := ix.Path(0); p != "0x0001/0x000000000000002a/0x00000003.geul" {
		t.Fatal(p)
	}
}
