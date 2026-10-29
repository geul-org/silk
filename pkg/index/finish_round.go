//ff:func feature=index type=writer control=iteration dimension=1
//ff:what finishRound — apply a round's finished merges, save paused ones as checkpoints, commit a generation and collect old files; ErrPaused when any paused
package index

import (
	"errors"
	"fmt"
)

// finishRound writes compact.json before collecting, so paused outputs stay.
func finishRound(dir string, m *Manifest, outs []roundOut, rep *Report) error {
	var cps []checkpoint
	var errs []error
	applied := 0
	for _, o := range outs {
		rep.Written += o.res.Written
		if o.res.Checkpoint != nil {
			cps = append(cps, *o.res.Checkpoint)
			continue
		}
		if o.err != nil {
			errs = append(errs, o.err)
			continue
		}
		applyResult(m, o.res)
		applied++
		rep.Dropped += o.res.Dropped
		rep.Tasks = append(rep.Tasks, fmt.Sprintf("%s %s %d->1 (%d rows, %d dead dropped)", o.res.Task.Name, o.res.Task.Kind, len(o.res.Task.Inputs), o.res.Out.Rows, o.res.Dropped))
	}
	return errors.Join(saveRound(dir, m, cps, applied, rep), errors.Join(errs...))
}
