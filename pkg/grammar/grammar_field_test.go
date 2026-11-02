//ff:func feature=grammar type=scanner control=sequence
//ff:what TestGrammarField — the field cut, including a Rank word after an Object
package grammar

import (
	"bytes"
	"testing"
)

func TestGrammarField(t *testing.T) {
	g := testGrammar(t)
	tr := edgeOf(0x0A, tok("#1"), tok("P31"), tok("Q5"), tok("Q7"), []byte{0, 2})
	if v, ok := g.field(tr, 4); !ok || !v.Raw || !bytes.Equal(v.B, []byte{0, 2}) {
		t.Fatal(v, ok)
	}
	if _, ok := g.field(edgeOf(0x0A, tok("#1"), tok("P31")), 3); ok {
		t.Fatal("missing fields")
	}
	if _, ok := g.field(edgeOf(0x0A, tok("#1"), tok("P31")), 2); ok {
		t.Fatal("missing target field")
	}
	if _, ok := g.field(edgeOf(0x0A, tok("#1")), 1); ok {
		t.Fatal("empty")
	}
	if _, _, ok := func() (int, int, bool) { _, ok := g.cut(tr, 0x100A, g.Fields[0x100A], 3, 40); return 0, 0, ok }(); ok {
		t.Fatal("past the end")
	}
}
