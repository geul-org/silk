//ff:func feature=query type=scanner control=sequence
//ff:what TestQueryWord — a new condition, a predicate, a predicate too early, a bad TAG
package query

import (
	"errors"
	"testing"
)

func TestQueryWord(t *testing.T) {
	_, _, ix := smallIndex(t)
	var q Query
	if err := q.word(ix, "a=1"); !errors.Is(err, ErrSyntax) {
		t.Fatal(err)
	}
	if err := q.word(ix, "NOPE"); !errors.Is(err, ErrSyntax) {
		t.Fatal(err)
	}
	q = Query{}
	if err := q.word(ix, "TRIPLE"); err != nil || len(q.Conds) != 1 {
		t.Fatal(err)
	}
	if err := q.word(ix, "predicate=P31"); err != nil || len(q.Conds[0].Preds) != 1 {
		t.Fatal(err)
	}
}
