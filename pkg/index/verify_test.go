//ff:func feature=index type=rule control=sequence
//ff:what TestVerify — a built index verifies; missing catalog, moved sets, broken maps and extra catalog entries fail
package index

import (
	"context"
	"errors"
	"testing"

	"github.com/park-jun-woo/silk/pkg/archive"
)

func TestVerify(t *testing.T) {
	ctx := context.Background()
	a, dir := fixtureArchive(t)
	if _, err := Verify(ctx, a, dir); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
	Build(ctx, a, dir)
	if n, err := Verify(ctx, a, dir); err != nil || n != 1 {
		t.Fatal(n, err)
	}
	moved := &archive.Archive{Sets: append([]archive.Set{}, a.Sets...)}
	moved.Sets[0].Generation = 2
	if _, err := Verify(ctx, moved, dir); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
	c, _ := readCatalog(dir)
	c.Maps = append(c.Maps, "0x0009/x.srm")
	writeCatalog(dir, c)
	if _, err := Verify(ctx, a, dir); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
	c.Maps = c.Maps[:1]
	writeCatalog(dir, c)
	removeStale(dir, nil)
	if _, err := Verify(ctx, a, dir); err == nil {
		t.Fatal("missing map")
	}
}
