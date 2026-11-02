//ff:func feature=query type=scanner control=sequence
//ff:what TestNewCond — a section name, a TAG name, neither
package query

import (
	"errors"
	"testing"
)

func TestNewCond(t *testing.T) {
	_, _, ix := smallIndex(t)
	if c, err := newCond(ix, "entity"); err != nil || c.Decl == nil {
		t.Fatal(c, err)
	}
	if c, err := newCond(ix, "LABEL"); err != nil || c.Decl != nil || c.Tags[0] != 0x1004 {
		t.Fatal(c, err)
	}
	if _, err := newCond(ix, "NOPE"); !errors.Is(err, ErrSyntax) {
		t.Fatal(err)
	}
}
