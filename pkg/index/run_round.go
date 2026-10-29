//ff:func feature=index type=builder control=iteration dimension=1
//ff:what runRound — run a round's tasks (one per section) on at most Threads goroutines; outcomes in task order
package index

import (
	"context"
	"sync"
)

// runRound resumes a task whose output has a checkpoint in cps.
func runRound(ctx context.Context, dir string, ts []Task, cps map[string]*checkpoint, live Bitmap, o Options) []roundOut {
	th := o.Threads
	if th <= 0 {
		th = DefaultThreads
	}
	sem := make(chan struct{}, th)
	outs := make([]roundOut, len(ts))
	var wg sync.WaitGroup
	for i, t := range ts {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			outs[i].res, outs[i].err = mergeTask(ctx, dir, t, live.Get, cps[t.Output], o.Pause)
		}()
	}
	wg.Wait()
	return outs
}
