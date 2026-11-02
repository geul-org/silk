//ff:func feature=index type=builder control=iteration dimension=1
//ff:what compactor.runTasks — do the round's merges in turn and commit them as one generation
package index

import "context"

func (c *compactor) runTasks(ctx context.Context, ts []task, r *Report) error {
	for _, t := range ts {
		if err := c.merge(ctx, t); err != nil {
			return err
		}
		r.Merges++
	}
	if err := commit(c.dir, &c.m); err != nil {
		return err
	}
	collect(c.dir, c.m)
	return nil
}
