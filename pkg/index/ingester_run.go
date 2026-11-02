//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what ingester.run — feed workers the documents wanted, in number order
package index

import (
	"context"
	"errors"
	"sync"
)

func (g *ingester) run(ctx context.Context, docs *DocTable) error {
	jobs := make(chan job, 2*threadsOf(g.u.o))
	var wg sync.WaitGroup
	for range threadsOf(g.u.o) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			g.worker(ctx, jobs)
		}()
	}
	err := g.feed(ctx, docs, jobs)
	close(jobs)
	wg.Wait()
	return errors.Join(err, g.err)
}
