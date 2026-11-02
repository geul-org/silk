//ff:func feature=index type=writer control=iteration dimension=1
//ff:what ingester.sink — move a document's rows to the shared buffers; past the budget every buffer becomes a run
package index

import "github.com/geul-org/silk/pkg/seg"

func (g *ingester) sink(local [][]*seg.Buffer, size uint64) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.maxDoc = max(g.maxDoc, size)
	g.read++
	var used int64
	for i, t := range g.ts {
		for k, sp := range t.sp {
			sp.buf.Absorb(local[i][k])
			used += sp.buf.Size()
		}
	}
	if used < g.budget {
		return nil
	}
	return g.flushAll()
}
