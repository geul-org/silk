//ff:func feature=decl type=scanner control=sequence
//ff:what TestDeclInfos — row shapes of the four kinds; the hash follows the content; a bad kind and bad silk fail
package decl

import (
	"errors"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestDeclInfos(t *testing.T) {
	g := fixture.Grammar(t)
	ix, _ := NewIndex(g, "t", "TRIPLE", "predicate", []string{"object"}, []int{6})
	if is, _ := ix.Infos(); is[0].RowWidth != 14 || is[0].DocOff != 6 || is[0].PartWidth != 6 {
		t.Fatal(is)
	}
	tx, _ := NewText(g, "l", "LABEL")
	if is, _ := tx.Infos(); is[0].RowWidth != 16 || is[0].PartWidth != LangWidth {
		t.Fatal(is)
	}
	sl, _ := NewSilk("s", silkText)
	if is, _ := sl.Infos(); len(is) != 2 || is[0].RowWidth != 6 || is[0].DocOff != 2 || is[1].DocOff != 0 {
		t.Fatal(is)
	}
	cp, err := NewComposite(g, "c", "LABEL", semText)
	if err != nil {
		t.Fatal(err)
	}
	if is, _ := cp.Infos(); is[0].RowWidth != 20 || is[0].DocOff != 16 {
		t.Fatal(is)
	}
	if ix.Hash() == tx.Hash() {
		t.Fatal("hash")
	}
	if _, err := (Decl{Kind: "nope"}).Infos(); !errors.Is(err, ErrDecl) {
		t.Fatal(err)
	}
	if _, err := (Decl{Kind: KindSilk, Silk: "x"}).Infos(); err == nil {
		t.Fatal("bad silk")
	}
	if sum([]int{1, 2}) != 3 {
		t.Fatal("sum")
	}
}
