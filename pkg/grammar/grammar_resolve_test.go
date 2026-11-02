//ff:func feature=grammar type=scanner control=iteration dimension=1
//ff:what TestGrammarResolve — names and numbers reach the same field; unknown parts fail; Width follows the kind
package grammar

import (
	"errors"
	"slices"
	"testing"
)

func TestGrammarResolve(t *testing.T) {
	g := testGrammar(t)
	a, err := g.Resolve(0x1004, "object.bytes")
	b, err2 := g.Resolve(0x1004, "2.1")
	if err != nil || err2 != nil || !slices.Equal(a.Path, b.Path) || a.Width() != 32 {
		t.Fatal(a, b, err, err2)
	}
	if l, _ := g.Resolve(0x1004, "object.language"); l.Width() != 16 {
		t.Fatal(l)
	}
	if l, _ := g.Resolve(0x100A, "rank"); l.Width() != 2 {
		t.Fatal(l)
	}
	if l, _ := g.Resolve(0x100A, "predicate"); l.Width() != 6 {
		t.Fatal(l)
	}
	if l, err := g.Resolve(0x100A, "0"); err != nil || l.Field.Name != "id" {
		t.Fatal(l, err)
	}
	for _, s := range []string{"nope", "9", "object.nope", "rank.x"} {
		if _, err := g.Resolve(0x100A, s); !errors.Is(err, ErrSyntax) {
			t.Fatal(s, err)
		}
	}
	if _, _, _, err := g.step([]uint16{0x7777}, "x"); !errors.Is(err, ErrSyntax) {
		t.Fatal(err)
	}
	if _, _, ok := g.lookup(0x084A, "0"); !ok {
		t.Fatal("value field 0")
	}
	if _, _, ok := lookupName(nil, "x", 0); ok {
		t.Fatal("empty")
	}
}
