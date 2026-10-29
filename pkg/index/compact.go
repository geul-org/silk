//ff:func feature=index type=writer control=sequence
//ff:what Compact — take the writer lock and run the due merges of the current generation (resuming paused ones first)
package index

import (
	"context"
	"time"
)

// Compact merges until the policy finds nothing due (o.Full: until each
// section is one sorted run without dead rows). ErrPaused means a merge
// stopped at a checkpoint; calling Compact again resumes it.
func Compact(ctx context.Context, dir string, o Options) (Report, error) {
	t0 := time.Now()
	unlock, err := lockDir(dir)
	if err != nil {
		return Report{}, err
	}
	defer unlock()
	m, err := CurrentManifest(dir)
	if err != nil {
		return Report{}, err
	}
	rep, err := compact(ctx, dir, &m, o)
	rep.Elapsed = time.Since(t0)
	return rep, err
}
