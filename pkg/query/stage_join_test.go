//ff:func feature=query type=scanner control=sequence
//ff:what TestStageJoin — a later row condition joins by document; a condition with no rows leaves nothing
package query

import "testing"

func TestStageJoin(t *testing.T) {
	_, _, ix := smallIndex(t)
	if a := answer(t, ix, false, "TRIPLE", "predicate=P31", "object=Q5", "TRIPLE", "predicate=P106", "object=Q98"); len(a) != 0 {
		t.Fatal(a)
	}
	if a := answer(t, ix, false, "TRIPLE", "predicate=P31", "object=Q5", "TRIPLE", "predicate=P106", "object=Q169470"); len(a) != 2 {
		t.Fatal(a)
	}
	if a := answer(t, ix, false, "entity", "job=physicist", "LABEL", "object.language=@en", `object.bytes^="Albert"`); len(a) != 1 {
		t.Fatal(a)
	}
	if a := answer(t, ix, true, "entity", "job=physicist", "LABEL", "object.language=@en", `object.bytes^="Albert"`); len(a) != 1 {
		t.Fatal(a)
	}
	if a := answer(t, ix, true, "entity", "job=politician"); len(a) != 0 {
		t.Fatal(a)
	}
}
