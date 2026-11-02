//ff:func feature=query type=scanner control=sequence
//ff:what TestFillKeys — keys extend while predicates fix them; a prefix ends the run; a key without a predicate stops
package query

import (
	"testing"

	"github.com/geul-org/silk/pkg/decl"
)

func TestFillKeys(t *testing.T) {
	_, _, ix := smallIndex(t)
	q, _ := Parse(ix, []string{"LABEL", "object.language=@en", "object.bytes^=\"Ab\""})
	c := &q.Conds[0]
	p := &Plan{Cond: c, Decl: ix.Decls["label"], Prefix: [][]byte{nil}}
	used := fillPart(ix, p)
	fillKeys(ix, p, used)
	if len(p.Parts) != 2 || p.Fixed != 2 || string(p.Prefix[0]) != "Ab" || len(used) != 2 {
		t.Fatal(p)
	}
	d := decl.Decl{Tags: []uint16{0x1004}, Keys: []string{"object.bytes", "subject"}, Widths: []int{4, 6}}
	p = &Plan{Cond: c, Decl: d, Prefix: [][]byte{nil}}
	fillKeys(ix, p, map[int]int{})
	if p.Fixed != 1 {
		t.Fatal(p.Fixed)
	}
	p = &Plan{Cond: &Cond{}, Decl: d, Prefix: [][]byte{nil}}
	fillKeys(ix, p, map[int]int{})
	if p.Fixed != 0 {
		t.Fatal(p.Fixed)
	}
	s, _ := Parse(ix, []string{"LABEL", `subject=Q1`, "object.bytes=\"Ab\""})
	p = &Plan{Cond: &s.Conds[0], Decl: decl.Decl{Tags: []uint16{0x1004}, Keys: []string{"object.bytes", "subject"}, Widths: []int{4, 6}}, Prefix: [][]byte{nil}}
	fillKeys(ix, p, map[int]int{})
	if p.Fixed != 2 {
		t.Fatal(p.Fixed)
	}
}
