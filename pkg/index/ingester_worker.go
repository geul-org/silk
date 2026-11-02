//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what ingester.worker — read documents and hand their rows to the sink; the first error stops everyone
package index

import "context"

func (g *ingester) worker(ctx context.Context, jobs <-chan job) {
	local := g.locals()
	for j := range jobs {
		if g.failed.Load() {
			continue
		}
		if err := g.doc(ctx, j, local); err != nil {
			g.fail(err)
		}
	}
}
