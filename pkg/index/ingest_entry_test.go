//ff:func feature=index type=builder control=sequence
//ff:what TestIngestEntry — new packs and Issuers are registered once; the entry carries place, IDs and counts
package index

import (
	"testing"

	"github.com/geul-org/silk/pkg/archive"
)

func TestIngestEntry(t *testing.T) {
	m := &Manifest{}
	g := newIngest("d", m, smallSpec(), fill(7, true), 1)
	d := liveDoc{Set: "s", Pack: "p", Doc: archive.Doc{Issuer: "0x0001", ID: 5, Version: 2, Off: 8, Len: 10}}
	e1 := g.entry(d, []TagCount{{1, 1}})
	e2 := g.entry(d, nil)
	if len(m.Packs) != 1 || len(m.Issuers) != 1 || e1.Pack != 0 || e2.Issuer != 0 || e1.ID != 5 || e1.Len != 10 || e1.Off != 8 || e1.Version != 2 || len(e1.Counts) != 1 {
		t.Fatal(m, e1)
	}
}
