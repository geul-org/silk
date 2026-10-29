//ff:func feature=idle type=scanner control=iteration dimension=1
//ff:what Wait — block until the rule finds the system idle (polling), or the context ends; an unreadable system counts as idle
package idle

import (
	"context"
	"time"
)

// Wait polls sample every poll until r.Idle; a sample error (no /proc) does
// not block. It returns ctx.Err() when the context ends first.
func Wait(ctx context.Context, r Rule, poll time.Duration, sample func() (Sample, error)) error {
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		s, err := sample()
		if err != nil || r.Idle(s) {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(poll):
		}
	}
}
