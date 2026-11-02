//ff:func feature=query type=scanner control=sequence
//ff:what TestCompileBits — no mask, a bad definition, a bad condition, a good one
package query

import (
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/sig"
)

func TestCompileBits(t *testing.T) {
	if err := compileBits(&Plan{Decl: decl.Decl{Kind: decl.KindIndex}, Cond: &Cond{}}); err != nil {
		t.Fatal(err)
	}
	if err := compileBits(&Plan{Decl: decl.Decl{Kind: decl.KindComposite}, Cond: &Cond{}}); err != nil {
		t.Fatal(err)
	}
	p := &Plan{Decl: decl.Decl{Kind: decl.KindComposite, Silk: testSem}, Cond: &Cond{Silk: []sig.Cond{{Field: "type", Value: "person"}}}}
	if err := compileBits(p); err != nil || p.Bits == nil {
		t.Fatal(err)
	}
	p.Cond.Silk[0].Field = "nope"
	if err := compileBits(p); !errors.Is(err, sig.ErrCond) {
		t.Fatal(err)
	}
	p.Decl.Silk = "x"
	if err := compileBits(p); !errors.Is(err, sig.ErrDef) {
		t.Fatal(err)
	}
}
