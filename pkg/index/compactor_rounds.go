//ff:func feature=index type=builder control=iteration dimension=1
//ff:what compactor.rounds — one merge per section each round, committing after each round
package index

import "context"

func (c *compactor) rounds(ctx context.Context) (Report, error) {
	var r Report
	var err error
	if c.live, err = loadLiveOf(c.dir, c.m); err != nil {
		return r, err
	}
	for ctx.Err() == nil {
		ts, err := c.plan()
		if err != nil || len(ts) == 0 {
			return r, err
		}
		if err := c.runTasks(ctx, ts, &r); err != nil {
			return r, err
		}
		c.full = false
	}
	return r, ctx.Err()
}
