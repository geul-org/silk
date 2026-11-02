//ff:func feature=index type=scanner control=sequence
//ff:what TestCompactorRunTasks — a task's failure and a commit's stop the round; a good round commits
package index

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestCompactorRunTasks(t *testing.T) {
	ctx := context.Background()
	_, dir, _ := built(t, 2)
	o := Options{Limit: 1 << 21, Policy: DefaultPolicy()}
	c := &compactor{dir: dir, o: o, m: current(t, dir), tmp: t.TempDir()}
	s := &c.m.Sections[0]
	r := Report{}
	if err := c.runTasks(ctx, []task{{sec: s, runs: s.Runs}}, &r); err != nil || r.Merges != 1 {
		t.Fatal(err)
	}
	if err := c.runTasks(ctx, []task{{sec: s, runs: []SegRef{{File: "none"}}}}, &r); err == nil {
		t.Fatal("task")
	}
	file := filepath.Join(dir, "f")
	os.WriteFile(file, nil, 0o644)
	b := &compactor{dir: file, m: Manifest{}}
	if err := b.runTasks(ctx, nil, &r); err == nil {
		t.Fatal("commit")
	}
}
