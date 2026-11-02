//ff:func feature=decl type=scanner control=iteration dimension=1
//ff:what TestExtractorRows — each kind's rows from one document: slots, numbers, offsets, QIDs, spread and dedup
package decl

import (
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestExtractorRows(t *testing.T) {
	g := fixture.Grammar(t)
	doc := testDoc(t)
	ix, _ := NewIndex(g, "t", "TRIPLE", "predicate", []string{"object"}, []int{6})
	x, err := NewExtractor(ix, g)
	if err != nil || x.Unique() {
		t.Fatal(err)
	}
	b := buffers(t, ix)
	x.Rows(doc, 9, b)
	if b[0].Rows() != 3 {
		t.Fatal(b[0].Rows())
	}
	for _, name := range []string{"LABEL", "ALIAS", "BIND"} {
		d, _ := NewText(g, "x", name)
		x, _ := NewExtractor(d, g)
		b := buffers(t, d)
		x.Rows(doc, 9, b)
		if !x.Unique() {
			t.Fatal("unique")
		}
		b[0].Unique()
		want := map[string]int{"LABEL": 3, "ALIAS": 1, "BIND": 1}[name]
		if b[0].Rows() != want {
			t.Fatal(name, b[0].Rows())
		}
	}
	sl, _ := NewSilk("s", silkText)
	x, _ = NewExtractor(sl, g)
	bs := buffers(t, sl)
	x.Rows(doc, 9, bs)
	if bs[0].Rows() != 2 || bs[1].Rows() != 2 {
		t.Fatal(bs[0].Rows())
	}
	cp, _ := NewComposite(g, "c", "LABEL", semText)
	x, _ = NewExtractor(cp, g)
	bc := buffers(t, cp)
	x.Rows(doc, 9, bc)
	if bc[0].Rows() != 3 {
		t.Fatal(bc[0].Rows())
	}
	x.Rows(doc, 9, bc)
	if bc[0].Rows() != 6 {
		t.Fatal(bc[0].Rows())
	}
}
