//ff:func feature=index type=builder control=iteration dimension=1
//ff:what ingester.locals — one set of per-document buffers for a worker
package index

import "github.com/geul-org/silk/pkg/seg"

func (g *ingester) locals() [][]*seg.Buffer {
	out := make([][]*seg.Buffer, len(g.ts))
	for i, t := range g.ts {
		for _, sp := range t.sp {
			out[i] = append(out[i], seg.NewBuffer(sp.info))
		}
	}
	return out
}
