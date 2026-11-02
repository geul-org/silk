//ff:func feature=query type=scanner control=sequence
//ff:what TestPlanning — the longest key wins, ties go to the smaller section, no section refuses and suggests; a forced section is obeyed
package query

import (
	"errors"
	"strings"
	"testing"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/sig"
)

func TestPlanning(t *testing.T) {
	_, _, ix := smallIndex(t)
	plan := func(force string, w ...string) (*Plan, error) {
		q, err := Parse(ix, w)
		if err != nil {
			t.Fatal(err)
		}
		return makePlan(ix, &q.Conds[0], force)
	}
	p, err := plan("", "TRIPLE", "predicate=P106", "object=Q98")
	if err != nil || p.Name != "triple" || !p.Covered || p.Fixed != 2 {
		t.Fatal(p, err)
	}
	if !strings.Contains(p.describe(), "covered=true") {
		t.Fatal(p.describe())
	}
	p, err = plan("subject", "TRIPLE", "predicate=P106", "object=Q98")
	if err != nil || p.Name != "subject" || p.Covered {
		t.Fatal(p, err)
	}
	p, err = plan("", "LABEL", `object.bytes="Short"`)
	if err != nil || p.Name != "label" || len(p.Parts) != 0 {
		t.Fatal(p, err)
	}
	p, err = plan("", "LABEL", "object.language=@en", `object.bytes="Short"`)
	if err != nil || len(p.Parts) != 2 || p.Fixed != 2 {
		t.Fatal(p, err)
	}
	_, err = plan("", "DESCRIPTION", "subject=Q1")
	if !errors.Is(err, ErrNoPlan) || !strings.Contains(err.Error(), "create index") {
		t.Fatal(err)
	}
	if !better(ix, &Plan{Name: "a", Fixed: 2}, &Plan{Name: "b", Fixed: 1}) || better(ix, &Plan{Name: "t", Fixed: 1}, &Plan{Name: "triple", Fixed: 2}) {
		t.Fatal("fixed fields decide")
	}
	if !better(ix, &Plan{Name: "label", Fixed: 1}, nil) || better(ix, &Plan{Name: "label-sem"}, &Plan{Name: "label"}) {
		t.Fatal("ties")
	}
	if !better(ix, &Plan{Name: "alias"}, &Plan{Name: "label"}) {
		t.Fatal("fewer rows")
	}
	if suggest(&Cond{Name: "X"}) == "" {
		t.Fatal("suggest")
	}
	c := Cond{Decl: &decl.Decl{Silk: "x", Kind: decl.KindSilk}}
	if _, err := silkPlan(&c); err == nil {
		t.Fatal("bad silk")
	}
	c = Cond{Decl: &decl.Decl{Silk: testDef, Kind: decl.KindSilk}, Silk: []sig.Cond{{Field: "nope"}}}
	if _, err := silkPlan(&c); err == nil {
		t.Fatal("bad field")
	}
	if err := compileBits(&Plan{Decl: decl.Decl{Kind: decl.KindComposite, Silk: "x"}, Cond: &Cond{Silk: []sig.Cond{{}}}}); err == nil {
		t.Fatal("bad composite silk")
	}
	if p := candidate(ix, &Cond{Tags: []uint16{1}}, "triple", ""); p != nil {
		t.Fatal(p)
	}
	if !strings.Contains((&Plan{Pats: []sig.Pattern{{}}, Name: "s"}).describe(), "silk") {
		t.Fatal("silk describe")
	}
}
