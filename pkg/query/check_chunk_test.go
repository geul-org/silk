//ff:func feature=query type=matcher control=sequence
//ff:what TestCheckChunk — Edge Ids and match flags filled; an unreadable candidate and a record of another kind fail
package query

import (
	"errors"
	"testing"
)

func TestCheckChunk(t *testing.T) {
	ix := testIndex(t, 1)
	c := condOf(t, "P31=Q5")
	p, _ := planCond(ix, c)
	var st Stats
	cs, _ := selectCond(ix, p, &st)
	ids, ok := make([]uint64, len(cs)), make([]bool, len(cs))
	if err := checkChunk(ix, cs, c, ids, ok); err != nil || !ok[0] || ids[0] != 15 {
		t.Fatal(err, ids, ok)
	}
	if err := checkChunk(ix, []candidate{{doc: 0, off: 1}}, c, ids, ok); err == nil {
		t.Fatal("unreadable")
	}
	if err := checkChunk(ix, []candidate{{doc: 0, off: 0}}, c, ids, ok); !errors.Is(err, ErrVerify) {
		t.Fatal(err)
	}
}
