//ff:func feature=index type=helper control=sequence
//ff:what TestEachIssuer — one target per Issuer in order with its documents; an error stops
package index

import (
	"errors"
	"testing"

	"github.com/park-jun-woo/silk/pkg/archive"
)

func TestEachIssuer(t *testing.T) {
	p := archive.Pack{Docs: []archive.Doc{{Issuer: "0x0002", ID: 1}, {Issuer: "0x0001", ID: 2}, {Issuer: "0x0002", ID: 3}}}
	var got []Target
	if err := eachIssuer(p, func(x Target) error { got = append(got, x); return nil }); err != nil || len(got) != 2 || got[0].Issuer != "0x0001" || len(got[1].Docs) != 2 {
		t.Fatal(got, err)
	}
	stop := errors.New("stop")
	if err := eachIssuer(p, func(Target) error { return stop }); err != stop {
		t.Fatal(err)
	}
}
