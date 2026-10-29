//ff:func feature=index type=helper control=sequence
//ff:what TestByIssuer — Issuers sorted, each group in input order
package index

import (
	"slices"
	"testing"

	"github.com/park-jun-woo/silk/pkg/archive"
)

func TestByIssuer(t *testing.T) {
	names, by := byIssuer([]archive.Doc{{Issuer: "0x0002", ID: 1}, {Issuer: "0x0001", ID: 2}, {Issuer: "0x0002", ID: 3}})
	if !slices.Equal(names, []string{"0x0001", "0x0002"}) || len(by["0x0002"]) != 2 || by["0x0002"][1].ID != 3 {
		t.Fatal(names, by)
	}
}
