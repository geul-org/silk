//ff:func feature=decl type=scanner control=sequence
//ff:what TestSubjectID — a record without a Subject gives 0; the language slot is empty without a language
package decl

import (
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/grammar"
)

func TestSubjectID(t *testing.T) {
	g := fixture.Grammar(t)
	d, _ := NewText(g, "x", "BIND")
	x, _ := NewExtractor(d, g)
	doc := &grammar.Doc{B: []byte{0x10, 0x03}}
	if n := x.subjectID(doc, grammar.Rec{Off: 0, Len: 2}); n != 0 {
		t.Fatal(n)
	}
	if s := x.langSlot(nil); len(s) != 0 {
		t.Fatal(s)
	}
}
