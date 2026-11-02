//ff:func feature=decl type=scanner control=sequence
//ff:what TestNewText — LABEL partitions by language, BIND does not; a pair or a non-text Edge is refused
package decl

import (
	"errors"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestNewText(t *testing.T) {
	g := fixture.Grammar(t)
	d, err := NewText(g, "b", "BIND")
	if err != nil || d.Part != "" || d.Keys[0] != "name.bytes" {
		t.Fatal(d, err)
	}
	if _, err := textBase(g, "x", "NOPE"); err == nil {
		t.Fatal("unknown")
	}
	if _, err := textBase(g, "x", "UTF8-TEXT"); !errors.Is(err, ErrDecl) {
		t.Fatal("pair")
	}
	if _, err := textBase(g, "x", "TRIPLE"); !errors.Is(err, ErrDecl) {
		t.Fatal("not text")
	}
}
