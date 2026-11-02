//ff:func feature=decl type=scanner control=sequence
//ff:what TestWithWidths — given widths, defaults and refusals
package decl

import (
	"errors"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestWithWidths(t *testing.T) {
	g := fixture.Grammar(t)
	if d, err := withWidths(Decl{}, g, 0x1004, []string{"subject"}, []int{4}); err != nil || d.Widths[0] != 4 {
		t.Fatal(d, err)
	}
	if d, err := withWidths(Decl{}, g, 0x1004, []string{"subject"}, nil); err != nil || d.Widths[0] != 6 {
		t.Fatal(d, err)
	}
	if _, err := withWidths(Decl{}, g, 0x1004, []string{"subject"}, []int{-1}); !errors.Is(err, ErrDecl) {
		t.Fatal(err)
	}
	if _, err := withWidths(Decl{}, g, 0x1004, []string{"nope"}, nil); err == nil {
		t.Fatal("path")
	}
	if _, err := withWidths(Decl{}, g, 0x1004, []string{"subject"}, []int{1, 2}); !errors.Is(err, ErrDecl) {
		t.Fatal(err)
	}
}
