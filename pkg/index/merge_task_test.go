//ff:func feature=index type=builder control=sequence
//ff:what TestMergeTask — a major merge drops the dead rows; a paused merge resumes to the same file; missing inputs, a blocked output and a writer error fail
package index

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestMergeTask(t *testing.T) {
	_, _, dir := updated(t)
	task, live, m := majorTask(t, dir)
	ctx := context.Background()
	r, err := mergeTask(ctx, dir, task, live.Get, nil, nil)
	if err != nil || r.Dropped != m.Sections[0].Dead || r.Out.Rows != rowsOf(task.Inputs)-r.Dropped || r.Written != r.Out.Bytes {
		t.Fatal(err, r)
	}
	whole, _ := os.ReadFile(filepath.Join(dir, task.Output))
	os.Remove(filepath.Join(dir, task.Output))
	n := 0
	pause := func(context.Context) error { n++; return map[bool]error{true: io.EOF, false: nil}[n == 2] }
	p, err := mergeTask(ctx, dir, task, live.Get, nil, pause)
	if !errors.Is(err, ErrPaused) || p.Checkpoint == nil || p.Checkpoint.Last == nil {
		t.Fatal(err, p)
	}
	r2, err := mergeTask(ctx, dir, task, live.Get, p.Checkpoint, nil)
	again, _ := os.ReadFile(filepath.Join(dir, task.Output))
	if err != nil || r2.Dropped != r.Dropped || !bytes.Equal(whole, again) {
		t.Fatal(err, r2)
	}
	cctx, cancel := context.WithCancel(ctx)
	cancel()
	if _, err := mergeTask(cctx, dir, task, live.Get, nil, nil); !errors.Is(err, ErrPaused) {
		t.Fatal("cancelled context", err)
	}
	bad := task
	bad.Inputs = append(bad.Inputs, SegRef{File: "seg/none"})
	if _, err := mergeTask(ctx, dir, bad, live.Get, nil, nil); err == nil {
		t.Fatal("missing input")
	}
	block(dir, "out")
	bad = task
	bad.Output = "out/x.sseg"
	if _, err := mergeTask(ctx, dir, bad, live.Get, nil, nil); err == nil {
		t.Fatal("blocked output")
	}
	p2, _ := mergeTask(ctx, dir, task, live.Get, nil, func(context.Context) error { return io.EOF })
	p2.Checkpoint.Last, p2.Checkpoint.State.Key = nil, bytes.Repeat([]byte{0xff}, 6)
	if _, err := mergeTask(ctx, dir, task, live.Get, p2.Checkpoint, nil); err == nil || errors.Is(err, ErrPaused) {
		t.Fatal("writer error", err)
	}
	if _, err := os.Stat(filepath.Join(dir, task.Output+".tmp")); !os.IsNotExist(err) {
		t.Fatal("aborted output kept")
	}
}
