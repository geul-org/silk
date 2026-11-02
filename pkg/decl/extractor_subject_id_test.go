//ff:func feature=decl type=scanner control=iteration dimension=1
//ff:what TestExtractorSubjectID — a Subject's QID word, and 0 without one
package decl

import (
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/grammar"
)

func TestExtractorSubjectID(t *testing.T) {
	g := fixture.Grammar(t)
	d, _ := NewText(g, "x", "BIND")
	x, _ := NewExtractor(d, g)
	doc := testDoc(t)
	var rec grammar.Rec
	for _, r := range doc.Recs {
		if r.Tag == 0x1003 {
			rec = r
		}
	}
	if n := x.subjectID(doc, rec); n != 0x40000001 {
		t.Fatalf("%x", n)
	}
	y := &grammar.Doc{B: []byte{0x10, 0x03}}
	if n := x.subjectID(y, grammar.Rec{Len: 2}); n != 0 {
		t.Fatal(n)
	}
}
