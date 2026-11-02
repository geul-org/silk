//ff:func feature=decl type=scanner control=sequence
//ff:what TestNewSilk — a bad definition is refused; a composite needs 32 bits
package decl

import (
	"errors"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestNewSilk(t *testing.T) {
	g := fixture.Grammar(t)
	if _, err := NewSilk("s", "x"); err == nil {
		t.Fatal("bad")
	}
	if _, err := NewComposite(g, "c", "LABEL", "x"); err == nil {
		t.Fatal("bad")
	}
	if _, err := NewComposite(g, "c", "LABEL", silkText); !errors.Is(err, ErrDecl) {
		t.Fatal(err)
	}
}
