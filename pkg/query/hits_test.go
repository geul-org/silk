//ff:func feature=query type=scanner control=sequence
//ff:what TestHits — joining, sorting, bare hits and unique documents
package query

import (
	"testing"

	"github.com/geul-org/silk/pkg/decl"
)

func TestHits(t *testing.T) {
	a := []Hit{{Doc: 1, Off: 2, HasOff: true}, {Doc: 2}}
	b := []Hit{{Doc: 2, Off: 1, HasOff: true}, {Doc: 3}}
	j := joinDocs(a, b)
	if len(j) != 2 || j[0].Doc != 2 {
		t.Fatal(j)
	}
	if s := sortHits([]Hit{{Doc: 2}, {Doc: 1}, {Doc: 1}}); len(s) != 2 || s[0].Doc != 1 {
		t.Fatal(s)
	}
	if u := uniqueDocs([]Hit{{Doc: 2}, {Doc: 1}, {Doc: 2}}); len(u) != 2 {
		t.Fatal(u)
	}
	if h := trimBare([]Hit{{Doc: 1}, {Doc: 1, HasOff: true}, {Doc: 2}}); len(h) != 2 {
		t.Fatal(h)
	}
	p := &Plan{Decl: decl.Decl{Kind: decl.KindSilk}}
	if h := markLoose(p, []Hit{{}}); h[0].Loose {
		t.Fatal("silk rows keep their mark")
	}
	p.Decl.Kind = decl.KindText
	if h := markLoose(p, []Hit{{}}); !h[0].Loose {
		t.Fatal("text key")
	}
}
