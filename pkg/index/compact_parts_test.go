//ff:func feature=index type=scanner control=sequence
//ff:what TestCompactParts — planning and merging stop on missing files; the pause hook stops a merge
package index

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestCompactParts(t *testing.T) {
	ctx := context.Background()
	_, dir, _ := built(t, 2)
	m := current(t, dir)
	o := Options{Limit: 1 << 21, Policy: DefaultPolicy()}
	c := &compactor{dir: dir, o: o, m: m, tmp: t.TempDir()}
	boom := errors.New("stop")
	c.o.Pause = func() error { return boom }
	if err := c.pause(ctx)(); !errors.Is(err, boom) {
		t.Fatal(err)
	}
	c.o.Pause = func() error { return nil }
	if err := c.pause(ctx)(); err != nil {
		t.Fatal(err)
	}
	s := &m.Sections[0]
	if t2, ok, err := c.planSection(&Section{}); ok || err != nil {
		t.Fatal(t2, ok, err)
	}
	if err := c.merge(ctx, task{sec: s, runs: []SegRef{{File: "none"}}}); err == nil {
		t.Fatal("missing input")
	}
	if _, _, err := c.planSection(&Section{Runs: []SegRef{{File: "none"}}}); err == nil {
		t.Fatal("missing input")
	}
	c.m.Sections = []Section{{Runs: []SegRef{{File: "none"}}}}
	if _, err := c.plan(); err == nil {
		t.Fatal("missing input")
	}
	if _, err := c.rounds(ctx); err == nil {
		t.Fatal("missing input")
	}
	info, _ := infoOf(filepath.Join(dir, s.Runs[0].File))
	if _, _, err := c.mergeInputs([]string{filepath.Join(dir, "none")}, info, seg.Hooks{}); err == nil {
		t.Fatal("missing input")
	}
	if _, err := c.place(filepath.Join(dir, "none"), 1); err == nil {
		t.Fatal("missing output")
	}
	bad := &compactor{dir: dir, m: Manifest{Live: "live/none"}}
	if _, err := bad.rounds(ctx); err == nil {
		t.Fatal("missing bitmap")
	}
	file := filepath.Join(dir, "f")
	os.WriteFile(file, nil, 0o644)
	nc := &compactor{dir: dir, o: o, m: m, tmp: file}
	if _, _, err := nc.mergeInputs([]string{filepath.Join(dir, s.Runs[0].File), filepath.Join(dir, s.Runs[0].File), filepath.Join(dir, s.Runs[0].File)}, info, seg.Hooks{}); err == nil {
		t.Fatal("tmp under a file")
	}
	ct := &compactor{dir: file, o: o, m: m, tmp: t.TempDir()}
	if err := ct.runTasks(ctx, nil, &Report{}); err == nil {
		t.Fatal("commit under a file")
	}
	if err := ct.runTasks(ctx, []task{{sec: s, runs: []SegRef{{File: "none"}}}}, &Report{}); err == nil {
		t.Fatal("task failure")
	}
	dp := filepath.Join(dir, "dead")
	w, _ := seg.NewWriter(dp, info)
	w.Finish()
	cm := &compactor{dir: dir, o: o, tmp: filepath.Join(file, "x")}
	if err := cm.merge(ctx, task{sec: &Section{}, runs: []SegRef{{File: "dead"}}}); err == nil {
		t.Fatal("no tmp directory")
	}
}
