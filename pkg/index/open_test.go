//ff:func feature=index type=reader control=sequence
//ff:what TestOpen — a built index opens with one part; no catalog, moved sets and a broken listed map fail
package index

import (
	"context"
	"errors"
	"testing"

	"github.com/park-jun-woo/silk/pkg/archive"
)

func TestOpen(t *testing.T) {
	a, dir := fixtureArchive(t)
	if _, err := Open(a, dir); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
	Build(context.Background(), a, dir)
	ix, err := Open(a, dir)
	if err != nil || len(ix.Parts) != 1 || ix.Parts[0].Map.Header.Docs != 4 {
		t.Fatal(ix, err)
	}
	ix.Close()
	moved := &archive.Archive{Sets: []archive.Set{{Name: "0x00010000", Generation: 2}}}
	if _, err := Open(moved, dir); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
	c, _ := readCatalog(dir)
	c.Maps = append(c.Maps, "0x0009/x.srm")
	writeCatalog(dir, c)
	if _, err := Open(a, dir); err == nil {
		t.Fatal("missing listed map")
	}
}
