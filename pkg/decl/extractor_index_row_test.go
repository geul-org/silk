//ff:func feature=decl type=scanner control=iteration dimension=1
//ff:what TestExtractorIndexRow — slots from a record: a partition and keys, missing fields zero
package decl

import (
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/grammar"
)

func TestExtractorIndexRow(t *testing.T) {
	g := fixture.Grammar(t)
	d, _ := NewIndex(g, "t", "LABEL", "object.language", []string{"subject"}, nil)
	x, _ := NewExtractor(d, g)
	doc := testDoc(t)
	var r grammar.Rec
	for _, c := range doc.Recs {
		if c.Tag == 0x1004 {
			r = c
			break
		}
	}
	p, k := x.indexRow(doc, doc.At(r), nil, nil)
	if len(p) != 16 || len(k) != 6 || p[1] != 0xC6 {
		t.Fatal(p, k)
	}
	u, _ := NewIndex(g, "u", "LABEL", "", []string{"subject"}, nil)
	y, _ := NewExtractor(u, g)
	if p, _ := y.indexRow(doc, doc.At(r), nil, nil); len(p) != 0 {
		t.Fatal(p)
	}
}
