//ff:func feature=query type=scanner control=iteration dimension=1
//ff:what TestRunnerMatches — every predicate holds; a missing field or another value fails
package query

import (
	"testing"

	"github.com/geul-org/silk/pkg/grammar"
)

func TestRunnerMatches(t *testing.T) {
	_, _, ix := smallIndex(t)
	r := testRunner(ix)
	d, _ := r.doc(1)
	rec := d.At(d.Recs[1])
	for i := range d.Recs {
		if d.Recs[i].Tag == 0x100A {
			rec = d.At(d.Recs[i])
			break
		}
	}
	yes, _ := Parse(ix, []string{"TRIPLE", "predicate=P31", "object=Q99"})
	no, _ := Parse(ix, []string{"TRIPLE", "predicate=P31", "object=Q98"})
	if !r.matches(d, rec, &yes.Conds[0]) || r.matches(d, rec, &no.Conds[0]) {
		t.Fatal("matches")
	}
	far := &Cond{Preds: []Pred{{Leaf: grammar.Leaf{Path: grammar.Path{9}}}}}
	if r.matches(d, rec, far) {
		t.Fatal("no such field")
	}
}
