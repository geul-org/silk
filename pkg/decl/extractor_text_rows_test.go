//ff:func feature=decl type=scanner control=sequence
//ff:what TestExtractorTextRows — a record without text gives no row
package decl

import (
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/grammar"
)

func TestExtractorTextRows(t *testing.T) {
	g := fixture.Grammar(t)
	d, _ := NewText(g, "t", "LABEL")
	if b := rowsOfKind(t, d); b[0].Rows() != 3 {
		t.Fatal(b[0].Rows())
	}
	x, _ := NewExtractor(d, g)
	doc, _ := grammar.Frame([]byte{0x11, 0x04, 0, 0x88, 0, 1, 0, 0xA0, 0, 1})
	b := buffers(t, d)[0]
	x.textRows(doc, 1, b)
	if b.Rows() != 0 {
		t.Fatal(b.Rows())
	}
}
