//ff:func feature=index type=writer control=sequence
//ff:what finishWriter — after a build or update, run compaction when asked and add its work to the report
package index

import (
	"context"
	"time"
)

// finishWriter compacts with o (when o.Full or o.Compact) under the held lock.
func finishWriter(ctx context.Context, dir string, m *Manifest, o Options, rep Report, t0 time.Time) (Report, error) {
	var err error
	if o.Full || o.Compact {
		var c Report
		c, err = compact(ctx, dir, m, o)
		rep.Written += c.Written
		rep.Dropped += c.Dropped
		rep.Tasks = append(rep.Tasks, c.Tasks...)
		rep.Paused = c.Paused
		rep.Generation = m.Generation
	}
	rep.Elapsed = time.Since(t0)
	return rep, err
}
