//ff:func feature=query type=scanner control=sequence
//ff:what TestKeySlots — slots and coverage: rest fields, prefixes, values too wide for their slot
package query

import (
	"bytes"
	"testing"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/grammar"
)

func TestKeySlots(t *testing.T) {
	if b := appendKey(nil, Pred{Prefix: true}, true, []byte("abc"), 2); string(b) != "ab" {
		t.Fatal(b)
	}
	if b := appendKey(nil, Pred{}, false, []byte("abc"), 8); len(b) != 8 {
		t.Fatal(b)
	}
	if b := appendKey(nil, Pred{}, true, []byte("abcd"), 4); !bytes.Equal(b, []byte{'a', 'b', 'c', 0xFF}) {
		t.Fatal(b)
	}
	if b := extend([][]byte{{1}, {2}}, Pred{Vals: [][]byte{{3}, {4}}}, false, 1); len(b) != 4 {
		t.Fatal(b)
	}
	rest := grammar.Leaf{Field: grammar.Field{Width: grammar.RestWidth}}
	if !flagged(decl.Decl{Kind: decl.KindIndex}, rest) || flagged(decl.Decl{Kind: decl.KindText}, rest) || flagged(decl.Decl{Kind: decl.KindIndex}, grammar.Leaf{}) {
		t.Fatal("flagged")
	}
	if exactSlot(Pred{Vals: [][]byte{{1, 2, 3}}}, 2) || !exactSlot(Pred{Vals: [][]byte{{1}}}, 2) {
		t.Fatal("exactSlot")
	}
	p := &Plan{Decl: decl.Decl{Kind: decl.KindIndex}}
	c := &Cond{Preds: []Pred{{Text: "a=1", Vals: [][]byte{{1, 2, 3}}}}}
	if isCovered(c, map[int]int{0: 2}, p) {
		t.Fatal("too wide")
	}
	c.Preds[0].Vals[0] = []byte{1}
	if !isCovered(c, map[int]int{0: 2}, p) {
		t.Fatal("fits")
	}
	if isCovered(c, map[int]int{}, p) {
		t.Fatal("unused")
	}
	c.Preds[0].Prefix = true
	if isCovered(c, map[int]int{0: 2}, p) {
		t.Fatal("prefix")
	}
	if predAt(c, grammar.Path{9}) != -1 {
		t.Fatal("predAt")
	}
}
