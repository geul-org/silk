//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what ingester.feed — queue live documents: new ones for everybody, old ones when a pending section needs them
package index

import "context"

func (g *ingester) feed(ctx context.Context, docs *DocTable, jobs chan<- job) error {
	from := g.u.first
	if g.anyPending() {
		from = 0
	}
	for seq := from; seq < g.u.m.DocCount() && !g.failed.Load(); seq++ {
		if !g.u.live.Get(seq) {
			continue
		}
		select {
		case jobs <- job{seq: seq, entry: docs.At(seq), isNew: seq >= g.u.first}:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}
