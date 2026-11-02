//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what TestIngestErrors — a section whose declaration, grammar or document breaks stops the run with that error
package index

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestIngestErrors(t *testing.T) {
	ctx := context.Background()
	_, dir, a := built(t, 2)
	m := current(t, dir)
	u := &updater{ar: a, dir: dir, o: testOpts, m: m, tmp: t.TempDir()}
	u.loadLive()
	for i := range u.m.Sections {
		u.m.Sections[i].Pending = true
	}
	var r Report
	bad := u.m
	bad.Docs = []DocFile{{File: "doc/none", Count: 1}}
	v := copyOf(u)
	v.m = bad
	if err := v.ingest(ctx, &r); err == nil {
		t.Fatal("missing table")
	}
	v = copyOf(u)
	v.m.Grammar = "none"
	if err := v.ingest(ctx, &r); err == nil {
		t.Fatal("missing grammar")
	}
	if _, err := v.targets(); err == nil {
		t.Fatal("missing grammar")
	}
	v = copyOf(u)
	v.m.Sections = []Section{{Decl: "none"}}
	if _, err := v.targets(); err == nil {
		t.Fatal("missing declaration")
	}
	g, _ := loadGrammar(dir, m.Grammar)
	os.WriteFile(filepath.Join(dir, dirDecl, "broken.json"), []byte(`{"kind":"index","tags":[4106],"part":"nope"}`), 0o644)
	if _, err := v.newTarget(&Section{Decl: "broken"}, g); err == nil {
		t.Fatal("bad path")
	}
	os.WriteFile(filepath.Join(dir, dirDecl, "kind.json"), []byte(`{"kind":"nope"}`), 0o644)
	if _, err := v.newTarget(&Section{Decl: "kind"}, g); err == nil {
		t.Fatal("bad kind")
	}
	v = copyOf(u)
	v.m.Packs = []PackRef{{Set: "0x00010000", Hash: "gone"}, {Set: "0x00010000", Hash: "gone"}}
	if err := v.ingest(ctx, &r); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
	cctx, cancel := context.WithCancel(ctx)
	cancel()
	v = copyOf(u)
	v.o.Threads = 1
	if err := v.ingest(cctx, &r); err != nil && !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}
