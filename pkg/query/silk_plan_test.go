//ff:func feature=query type=scanner control=sequence
//ff:what TestSilkPlan — a good condition, a bad definition, a bad field
package query

import (
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/sig"
)

func TestSilkPlan(t *testing.T) {
	c := &Cond{Name: "e", Decl: &decl.Decl{Kind: decl.KindSilk, Silk: testDef}, Silk: []sig.Cond{{Field: "job", Value: "writer"}}}
	if p, err := silkPlan(c); err != nil || len(p.Pats) != 2 {
		t.Fatal(p, err)
	}
	c.Silk[0].Field = "nope"
	if _, err := silkPlan(c); !errors.Is(err, sig.ErrCond) {
		t.Fatal(err)
	}
	c.Decl.Silk = "x"
	if _, err := silkPlan(c); !errors.Is(err, sig.ErrDef) {
		t.Fatal(err)
	}
}
