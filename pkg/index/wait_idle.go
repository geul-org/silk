//ff:func feature=index type=helper control=sequence
//ff:what waitIdle — run the pause hook before a compaction round (nil without one)
package index

import "context"

// waitIdle lets the hook hold a round until the machine is idle.
func waitIdle(ctx context.Context, o Options) error {
	if o.Pause == nil {
		return nil
	}
	return o.Pause(ctx)
}
