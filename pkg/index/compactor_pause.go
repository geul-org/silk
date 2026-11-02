//ff:func feature=index type=helper control=sequence
//ff:what compactor.pause — the option's pause hook, also stopping on a cancelled context
package index

import "context"

func (c *compactor) pause(ctx context.Context) func() error {
	return func() error {
		if c.o.Pause != nil {
			if err := c.o.Pause(); err != nil {
				return err
			}
		}
		return ctx.Err()
	}
}
