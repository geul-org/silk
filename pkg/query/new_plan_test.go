//ff:func feature=query type=builder control=sequence
//ff:what TestNewPlan — four stages, the join only for several conditions, each condition planned; a bad condition fails
package query

import (
	"strings"
	"testing"
)

func TestNewPlan(t *testing.T) {
	ix := testIndex(t, 1)
	q := Query{Conds: []Cond{condOf(t, "P31=Q5"), condOf(t, "LABEL:2.0=@de")}}
	p, err := NewPlan(ix, q, 3)
	if err != nil || len(p.Steps) != 4 || len(p.Conds) != 2 || p.Steps[0].Active || !p.Steps[3].Active || !strings.Contains(p.Steps[1].Detail, "section LABEL") {
		t.Fatal(err, p)
	}
	if p1, _ := NewPlan(ix, Query{Conds: q.Conds[:1]}, 1); p1.Steps[3].Active {
		t.Fatal("join with one condition")
	}
	if _, err := NewPlan(ix, Query{Conds: []Cond{{Name: "X", Kind: 2, Codec: 3}}}, 1); err == nil {
		t.Fatal("bad condition")
	}
}
