//ff:func feature=query type=scanner control=sequence
//ff:what TestFillPart — none declared, none fixed, fixed
package query

import (
	"testing"

	"github.com/geul-org/silk/pkg/decl"
)

func TestFillPart(t *testing.T) {
	_, _, ix := smallIndex(t)
	q, _ := Parse(ix, []string{"LABEL", "object.language=@en"})
	if u := fillPart(ix, &Plan{Cond: &q.Conds[0], Decl: decl.Decl{}}); len(u) != 0 {
		t.Fatal(u)
	}
	if u := fillPart(ix, &Plan{Cond: &Cond{}, Decl: ix.Decls["label"]}); len(u) != 0 {
		t.Fatal(u)
	}
	p := &Plan{Cond: &q.Conds[0], Decl: ix.Decls["label"]}
	if u := fillPart(ix, p); len(u) != 1 || len(p.Parts) != 2 {
		t.Fatal(u, p.Parts)
	}
}
