//ff:func feature=index type=builder control=sequence
//ff:what mergeTask — run one merge: open the inputs, stream their groups into a new sorted run (dropping dead rows), pausing at partition boundaries; resumes from a checkpoint
package index

import (
	"context"
	"errors"

	"github.com/park-jun-woo/silk/pkg/seg"
)

// mergeTask returns ErrPaused with a checkpoint when ctx ends or pause fails.
func mergeTask(ctx context.Context, dir string, t Task, live func(uint32) bool, cp *checkpoint, pause func(context.Context) error) (taskResult, error) {
	fs, gs, err := openGroups(dir, t.Inputs)
	defer closeSegs([][]*seg.File{fs})
	if err != nil {
		return taskResult{}, err
	}
	w, from, base, err := openOutput(dir, t, fs[0].Info, cp)
	if err != nil {
		return taskResult{}, err
	}
	last := from
	after := func(key []byte) error {
		last = key
		if ctx.Err() != nil || (pause != nil && pause(ctx) != nil) {
			return ErrPaused
		}
		return nil
	}
	dropped, err := seg.Merge(gs, w, live, from, after)
	if cp != nil {
		dropped += cp.Dropped
	}
	if errors.Is(err, ErrPaused) {
		st, err := w.Suspend()
		c := &checkpoint{Section: t.Name, Kind: t.Kind, Inputs: refFiles(t.Inputs), Output: t.Output, Last: last, Dropped: dropped, State: st}
		return taskResult{Task: t, Written: w.Written() - base, Checkpoint: c}, errors.Join(ErrPaused, err)
	}
	if err != nil {
		w.Abort()
		return taskResult{}, err
	}
	info, err := w.Finish()
	return taskResult{Task: t, Out: SegRef{File: t.Output, Rows: info.Rows + info.Other, Bytes: info.Bytes}, Dropped: dropped, Written: info.Bytes - base}, err
}
