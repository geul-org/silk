//ff:func feature=index type=builder control=sequence
//ff:what TestBuildOne — one map written with the target's identity; missing pack, bad documents and write failures surface
package index

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/silk/pkg/archive"
	"github.com/park-jun-woo/silk/pkg/recmap"
)

func TestBuildOne(t *testing.T) {
	a, dir := fixtureArchive(t)
	ps, _ := a.Sets[0].Packs(context.Background())
	tg := Target{Pack: ps[0], Issuer: "0x00010000", Docs: ps[0].Docs}
	p := filepath.Join(dir, "m.srm")
	h, err := buildOne(p, tg)
	if err != nil || h.Pack != ps[0].Hash || h.Manifest != a.Sets[0].Manifest || h.Generation != 1 {
		t.Fatal(h, err)
	}
	if f, err := recmap.Open(p); err != nil || f.Header != h {
		t.Fatal(err)
	} else {
		f.Close()
	}
	missing := tg
	missing.Pack.Path = filepath.Join(dir, "none.gpk")
	if _, err := buildOne(p, missing); err == nil {
		t.Fatal("missing pack")
	}
	bad := tg
	bad.Docs = []archive.Doc{{Off: 0, Len: 64}}
	if _, err := buildOne(p, bad); err == nil {
		t.Fatal("bad document")
	}
	os.WriteFile(filepath.Join(dir, "file"), nil, 0o644)
	if _, err := buildOne(filepath.Join(dir, "file", "m.srm"), tg); err == nil {
		t.Fatal("write")
	}
}
