//ff:func feature=query type=builder control=sequence
//ff:what TestPlanCond — fields at the spec's partition and sort paths give slots (zero padded); other fields are left for the exact check; an unindexed kind or a field past the layout fails
package query

import (
	"bytes"
	"errors"
	"testing"
)

func TestPlanCond(t *testing.T) {
	ix := testIndex(t, 1)
	p, err := planCond(ix, condOf(t, "P106=Q169470"))
	if err != nil || !p.ByPart || !p.BySort || !bytes.Equal(p.Sel.Part, []byte{0, 0xa8, 0, 0x6a, 0, 0}) || len(p.Sel.Sort) != 6 {
		t.Fatalf("%v %x", err, p.Sel)
	}
	r, _ := planCond(ix, condOf(t, "REFERENCE-NODE:1=#189"))
	if r.ByPart || r.BySort || r.Section != 6 {
		t.Fatal(r)
	}
	if _, err := planCond(ix, Cond{Name: "BIND", Kind: 2, Codec: 3}); !errors.Is(err, ErrNoSection) {
		t.Fatal(err)
	}
	if _, err := planCond(ix, Cond{Name: "TRIPLE", Kind: 2, Codec: 0x0A, Fields: []FieldCond{{Path: []int{4}, Value: []byte{0, 0}}}}); !errors.Is(err, ErrSyntax) {
		t.Fatal(err)
	}
}
