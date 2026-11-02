//ff:func feature=index type=builder control=iteration dimension=1
//ff:what extract — one target's rows of one document
package index

import (
	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/seg"
)

func extract(t *target, d *grammar.Doc, seq uint32, out []*seg.Buffer) {
	t.x.Rows(d, seq, out)
	if !t.x.Unique() {
		return
	}
	for _, b := range out {
		b.Unique()
	}
}
