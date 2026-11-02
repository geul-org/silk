//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what TestUpdaterSyncDocs — new table and identity rows; a blocked table path and a missing section stop it
package index

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestUpdaterSyncDocs(t *testing.T) {
	ctx := context.Background()
	_, dir, a := setup(t, 2)
	mk := func(d string) *updater {
		return &updater{ar: a, dir: d, o: testOpts, m: Manifest{Paths: Section{Name: "paths"}}, tmp: t.TempDir()}
	}
	u := mk(dir)
	if err := u.syncDocs(ctx); err != nil || u.m.DocCount() != 2 || len(u.m.Issuers) != 1 {
		t.Fatal(err)
	}
	file := filepath.Join(dir, "f")
	os.WriteFile(file, nil, 0o644)
	if err := mk(file).syncDocs(ctx); err == nil {
		t.Fatal("table under a file")
	}
	v := mk(dir)
	v.m.Packs = []PackRef{{Set: "0x00010000", Hash: "gone"}}
	v.m.Docs = []DocFile{{File: "doc/none", Count: 1}}
	if err := v.syncDocs(ctx); err == nil {
		t.Fatal("killGone")
	}
	w := mk(dir)
	w.m.Paths.Runs = []SegRef{{File: "none"}}
	if err := w.syncDocs(ctx); err == nil {
		t.Fatal("settle")
	}
	x := mk(dir)
	x.m.Issuers = []string{"0x00010000"}
	os.MkdirAll(filepath.Join(dir, dirDoc, fmt.Sprintf("%016x.sdoc.tmp", 1)), 0o755)
	if err := x.syncDocs(ctx); err == nil {
		t.Fatal("table blocked")
	}
	os.RemoveAll(filepath.Join(dir, dirDoc))
	os.MkdirAll(filepath.Join(dir, dirDoc, fmt.Sprintf("%016x.sdoc", 1), "x"), 0o755)
	if err := mk(dir).syncDocs(ctx); err == nil {
		t.Fatal("rename blocked")
	}
	os.RemoveAll(filepath.Join(dir, dirDoc))
	st := a.Sets[0]
	es, _ := os.ReadDir(filepath.Join(st.Dir, "o"))
	for _, e := range es {
		os.Rename(filepath.Join(st.Dir, "o", e.Name()), filepath.Join(st.Dir, "o", e.Name()+"x"))
	}
	if err := mk(dir).syncDocs(ctx); err == nil {
		t.Fatal("replay")
	}
}
