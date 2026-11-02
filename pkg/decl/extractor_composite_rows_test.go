//ff:func feature=decl type=scanner control=sequence
//ff:what TestExtractorCompositeRows — a row per text Edge; a BIND-less record with no text is skipped
package decl

import (
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/grammar"
)

func TestExtractorCompositeRows(t *testing.T) {
	d, _ := NewComposite(fixture.Grammar(t), "c", "LABEL", semText)
	if b := rowsOfKind(t, d); b[0].Rows() != 3 {
		t.Fatal(b[0].Rows())
	}
	g := fixture.Grammar(t)
	x, _ := NewExtractor(d, g)
	doc, _ := grammar.Frame(grammar.RecordOf(fixture.Label(1, fixture.Tok("Q1"), "en", "x"))[:0])
	x.compositeRows(doc, 1, buffers(t, d)[0])
	bad := []byte{0x11, 0x04, 0, 0x88, 0, 1, 0, 0xA0, 0, 1}
	doc, err := grammar.Frame(bad)
	if err != nil {
		t.Fatal(err)
	}
	b := buffers(t, d)[0]
	x.compositeRows(doc, 1, b)
	if b.Rows() != 0 {
		t.Fatal(b.Rows())
	}
}
