//ff:func feature=index type=helper control=sequence
//ff:what TestEachPack — every pack visited in order; an error stops
package index

import (
	"errors"
	"testing"

	"github.com/park-jun-woo/silk/pkg/archive"
)

func TestEachPack(t *testing.T) {
	ps := []archive.Pack{{Hash: "a", Docs: []archive.Doc{{Issuer: "0x0001"}}}, {Hash: "b", Docs: []archive.Doc{{Issuer: "0x0001"}}}}
	var got []string
	if err := eachPack(ps, func(x Target) error { got = append(got, x.Pack.Hash); return nil }); err != nil || len(got) != 2 || got[1] != "b" {
		t.Fatal(got, err)
	}
	stop := errors.New("stop")
	if err := eachPack(ps, func(Target) error { return stop }); err != stop {
		t.Fatal(err)
	}
}
