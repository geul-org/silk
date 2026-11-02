//ff:func feature=decl type=scanner control=sequence
//ff:what TestTextBase — language text, plain text, no text, an unknown TAG
package decl

import (
	"errors"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestTextBase(t *testing.T) {
	g := fixture.Grammar(t)
	if d, err := textBase(g, "l", "LABEL"); err != nil || d.Part != "object.language" {
		t.Fatal(d, err)
	}
	if d, err := textBase(g, "b", "BIND"); err != nil || d.Part != "" || d.Widths[0] != TextWidth {
		t.Fatal(d, err)
	}
	if _, err := textBase(g, "x", "NOPE"); err == nil {
		t.Fatal("unknown")
	}
	if _, err := textBase(g, "x", "TRIPLE"); !errors.Is(err, ErrDecl) {
		t.Fatal("no text")
	}
	if _, err := textBase(g, "x", "UTF8-TEXT"); !errors.Is(err, ErrDecl) {
		t.Fatal("a pair")
	}
	if _, err := textBase(g, "x", "SOURCE"); !errors.Is(err, ErrDecl) {
		t.Fatal("no text at field 2")
	}
}
