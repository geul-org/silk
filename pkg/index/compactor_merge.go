//ff:func feature=index type=builder control=sequence
//ff:what compactor.merge — one task to one run without dead rows; the inputs leave the section
package index

import (
	"context"
	"os"

	"github.com/geul-org/silk/pkg/seg"
)

func (c *compactor) merge(ctx context.Context, t task) error {
	paths := segPaths(c.dir, Section{Runs: t.runs, Appends: t.appends})
	info, err := infoOf(paths[0])
	if err != nil {
		return err
	}
	hk := seg.Hooks{Live: c.live.Get, Pause: c.pause(ctx)}
	p, rows, err := c.mergeInputs(paths, info, hk)
	if err != nil {
		return err
	}
	t.sec.Runs = withoutSegs(t.sec.Runs, t.runs)
	t.sec.Appends = withoutSegs(t.sec.Appends, t.appends)
	if rows == 0 {
		return os.Remove(p)
	}
	ref, err := c.place(p, rows)
	t.sec.Runs = append(t.sec.Runs, ref)
	return err
}
