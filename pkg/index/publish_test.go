//ff:func feature=index type=scanner control=sequence
//ff:what TestPublish — small updates stay in the append array, and reaching AppendRows flushes appends and new rows into one sorted run
package index

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/seg"
)

func TestPublish(t *testing.T) {
	ctx := context.Background()
	data, dir, a := built(t, 2)
	o := testOpts
	o.Policy.AppendRows = 3
	fixture.Repack(t, data, docsOf(3), nil)
	a, _ = archive.Open(ctx, data)
	if _, err := Update(ctx, a, dir, o); err != nil {
		t.Fatal(err)
	}
	m := current(t, dir)
	if s := m.Sections[1]; len(s.Appends) != 1 || len(s.Runs) != 1 {
		t.Fatal(s)
	}
	fixture.Repack(t, data, docsOf(5), nil)
	a, _ = archive.Open(ctx, data)
	if _, err := Update(ctx, a, dir, o); err != nil {
		t.Fatal(err)
	}
	m = current(t, dir)
	if s := m.Sections[1]; len(s.Appends) != 0 || len(s.Runs) != 2 || s.Runs[1].Rows != 3 {
		t.Fatal(s)
	}
	if rowsOf([]SegRef{{Rows: 2}, {Rows: 3}}) != 5 {
		t.Fatal("rowsOf")
	}
	u := &updater{dir: dir, o: o, m: m, tmp: t.TempDir()}
	if _, err := u.place(filepath.Join(u.tmp, "none"), 1); err == nil {
		t.Fatal("missing")
	}
	if err := u.publish(&Section{Pending: true}, filepath.Join(u.tmp, "none"), 1); err == nil {
		t.Fatal("missing")
	}
	sec := Section{Appends: []SegRef{{File: "none"}}}
	if err := u.publishRun(&sec, filepath.Join(u.tmp, "none"), 1); err == nil {
		t.Fatal("missing")
	}
	p := filepath.Join(u.tmp, "r")
	w, _ := seg.NewWriter(p, pathInfo())
	w.Finish()
	if err := u.publishRun(&sec, p, 0); err == nil {
		t.Fatal("missing append")
	}
	os.WriteFile(filepath.Join(dir, dirSeg+"x"), nil, 0o644)
	file := filepath.Join(u.tmp, "f")
	os.WriteFile(file, nil, 0o644)
	bad := &updater{dir: file, m: m}
	if _, err := bad.place(p, 0); err == nil {
		t.Fatal("seg under a file")
	}
}
