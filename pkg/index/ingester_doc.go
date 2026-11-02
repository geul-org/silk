//ff:func feature=index type=builder control=iteration dimension=1
//ff:what ingester.doc — one document: read, frame, extract for every target that wants it, sink
package index

import (
	"context"
	"fmt"

	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/seg"
)

func (g *ingester) doc(ctx context.Context, j job, local [][]*seg.Buffer) error {
	b, err := g.u.read(ctx, j.entry)
	if err != nil {
		return err
	}
	d, err := grammar.Frame(b)
	if err != nil {
		return fmt.Errorf("document %d: %w", j.seq, err)
	}
	for i, t := range g.ts {
		if t.pending || j.isNew {
			extract(t, d, j.seq, local[i])
		}
	}
	return g.sink(local, uint64(len(b)))
}
