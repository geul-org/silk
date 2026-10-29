//ff:type feature=index type=model
//ff:what Options — writer settings: threads, rows per in-memory batch, compaction after update, full compaction, policy, pause hook
package index

import "context"

// Options tune Build, Update and Compact; zero values take the defaults.
// Pause, when set, is called between partitions of a merge and may block
// (waiting for an idle machine); an error stops the merge at a checkpoint.
type Options struct {
	Threads int
	Batch   uint64
	Compact bool
	Full    bool
	Policy  Policy
	Pause   func(ctx context.Context) error
}
