//ff:func feature=index type=scanner control=sequence
//ff:what TestUpdaterAdvance — each stage's failure stops the run: bitmap, table, documents, commit
package index

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestUpdaterAdvance(t *testing.T) {
	ctx := context.Background()
	_, dir, a := built(t, 2)
	mk := func() *updater {
		return &updater{ar: a, dir: dir, o: testOpts, m: current(t, dir), tmp: t.TempDir()}
	}
	u := mk()
	if _, err := u.advance(ctx); err != nil {
		t.Fatal(err)
	}
	u = mk()
	u.m.Live = "live/none"
	if _, err := u.advance(ctx); err == nil {
		t.Fatal("bitmap")
	}
	u = mk()
	u.m.Packs = nil
	u.m.Docs = append(u.m.Docs, DocFile{File: "doc/none", First: u.m.DocCount(), Count: 1})
	if _, err := u.advance(ctx); err == nil {
		t.Fatal("table")
	}
	file := filepath.Join(dir, "f")
	os.WriteFile(file, nil, 0o644)
	u = mk()
	u.dir = file
	u.m = Manifest{Paths: Section{Name: "paths"}}
	if _, err := u.advance(ctx); err == nil {
		t.Fatal("replay under a file")
	}
	u = mk()
	u.m.Grammar = "none"
	if _, err := u.advance(ctx); err == nil {
		t.Fatal("grammar")
	}
}
