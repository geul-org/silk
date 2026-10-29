//ff:func feature=index type=builder control=sequence
//ff:what TestNewIngest — a bucket per tag with its TAG, widths and minimum; known packs and Issuers numbered; batch at least 1
package index

import (
	"testing"

	"github.com/geul-org/silk/pkg/spec"
)

func TestNewIngest(t *testing.T) {
	m := &Manifest{Packs: []PackRef{{Set: "s", Hash: "h"}}, Issuers: []string{"0x0001"}}
	sp := spec.Default()
	g := newIngest("d", m, sp, fill(7, false), 0)
	b := g.buckets[0]
	if len(g.buckets) != 7 || b.Info.Tag != 0x100A || b.Info.PartWidth != 6 || b.Min != 256 || g.batch != 1 || g.packIx["s/h"] != 0 || g.issIx["0x0001"] != 0 || g.table[0x100A] != 0 {
		t.Fatal(g)
	}
}
