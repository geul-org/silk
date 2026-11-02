//ff:func feature=query type=scanner control=sequence
//ff:what TestPredAt — a path found, a path not
package query

import (
	"testing"

	"github.com/geul-org/silk/pkg/grammar"
)

func TestPredAt(t *testing.T) {
	c := &Cond{Preds: []Pred{{Leaf: grammar.Leaf{Path: grammar.Path{1}}}, {Leaf: grammar.Leaf{Path: grammar.Path{2, 1}}}}}
	if predAt(c, grammar.Path{2, 1}) != 1 || predAt(c, grammar.Path{3}) != -1 {
		t.Fatal("predAt")
	}
}
