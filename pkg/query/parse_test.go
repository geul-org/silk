//ff:func feature=query type=scanner control=iteration dimension=1
//ff:what TestParse — words open conditions and fill them; bad words fail
package query

import (
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/decl"
)

func TestParse(t *testing.T) {
	_, _, ix := smallIndex(t)
	q, err := Parse(ix, []string{"TRIPLE", "predicate=P31", "object=Q5", "entity", "job=physicist"})
	if err != nil || len(q.Conds) != 2 || len(q.Conds[0].Preds) != 2 || len(q.Conds[1].Silk) != 1 || q.Mode != ModeDocs {
		t.Fatal(q, err)
	}
	q, err = Parse(ix, []string{"label-sem", "object.bytes^=\"Al\"", "type=person"})
	if err != nil || len(q.Conds[0].Preds) != 1 || !q.Conds[0].Preds[0].Prefix || len(q.Conds[0].Silk) != 1 {
		t.Fatal(q, err)
	}
	for _, w := range [][]string{nil, {"predicate=P31"}, {"NOPE"}, {"TRIPLE", "nope=1"}, {"TRIPLE", "object=nope"}, {"entity", "object.bytes"}} {
		if _, err := Parse(ix, w); err == nil {
			t.Fatal(w)
		}
	}
	if _, _, _, ok := splitPred(`"text"`); ok {
		t.Fatal("quoted")
	}
	if _, _, _, ok := splitPred(""); ok {
		t.Fatal("empty")
	}
	if isPath(ix, &Cond{}, "x") {
		t.Fatal("no tags")
	}
	c := Cond{Decl: &decl.Decl{Kind: decl.KindSilk, Silk: "x"}, Tags: []uint16{0x1004}}
	if err := addPred(ix, &c, "object", "Q1", false); !errors.Is(err, ErrSyntax) {
		t.Fatal(err)
	}
}
