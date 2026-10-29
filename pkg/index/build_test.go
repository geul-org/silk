//ff:func feature=index type=builder control=sequence
//ff:what TestBuild — maps, catalog and stale-map removal; directory, target, catalog and removal failures surface
package index

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/silk/pkg/archive"
)

func TestBuild(t *testing.T) {
	ctx := context.Background()
	a, dir := fixtureArchive(t)
	os.MkdirAll(filepath.Join(dir, "0x00010000"), 0o755)
	old := filepath.Join(dir, "0x00010000", "old.srm")
	os.WriteFile(old, nil, 0o644)
	hs, err := Build(ctx, a, dir)
	if err != nil || len(hs) != 1 || hs[0].Docs != 4 || hs[0].Issuer != "0x00010000" {
		t.Fatal(hs, err)
	}
	if _, err := os.Stat(old); !os.IsNotExist(err) {
		t.Fatal("stale map kept", err)
	}
	if c, err := readCatalog(dir); err != nil || len(c.Maps) != 1 {
		t.Fatal(c, err)
	}
	file := filepath.Join(t.TempDir(), "f")
	os.WriteFile(file, nil, 0o644)
	if _, err := Build(ctx, a, filepath.Join(file, "idx")); err == nil {
		t.Fatal("directory under a file")
	}
	bad := &archive.Archive{Sets: []archive.Set{{Dir: t.TempDir(), Manifest: "none"}}}
	if _, err := Build(ctx, bad, t.TempDir()); err == nil {
		t.Fatal("bad set")
	}
	d2 := t.TempDir()
	os.MkdirAll(filepath.Join(d2, catalogFile+".tmp"), 0o755)
	if _, err := Build(ctx, a, d2); err == nil {
		t.Fatal("catalog write")
	}
	d3 := t.TempDir()
	os.MkdirAll(filepath.Join(d3, "0x0009", "x.srm", "y"), 0o755)
	if _, err := Build(ctx, a, d3); err == nil {
		t.Fatal("stale removal")
	}
}
