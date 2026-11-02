//ff:func feature=sig type=scanner control=sequence
//ff:what TestEntityCodes — no statement, markers, listed, unlisted, buckets across a decade, flags
package sig

import (
	"slices"
	"testing"
)

func TestEntityCodes(t *testing.T) {
	d := testDef(t)
	job, _ := d.field("job")
	alive := d.common("alive")
	born, _ := d.field("born")
	if c := ent().codes(job); !slices.Equal(c, []uint64{CodeNone}) {
		t.Fatal(c)
	}
	if c := ent("P106", "?", "P106", "!", "P106", "Q36180", "P106", "Q1", "P106", "Q36180").codes(job); !slices.Equal(c, []uint64{CodeNoValue, CodeSomeValue, CodeOther, 5}) {
		t.Fatal(c)
	}
	if c := ent("P570", "?").codes(alive); !slices.Equal(c, []uint64{1}) {
		t.Fatal(c)
	}
	if c := ent("P570", "!").codes(alive); !slices.Equal(c, []uint64{0}) {
		t.Fatal(c)
	}
	e := ent()
	e.Objects[string(born.Pred)] = [][]byte{edtf("18XX"), edtf("19XX"), []byte{0, 1}}
	if c := e.codes(born); !slices.Equal(c, []uint64{CodeSomeValue, 5, 6}) {
		t.Fatal(c)
	}
	if c := born.bucketCodes(edtf("x")); !slices.Equal(c, []uint64{CodeOther}) {
		t.Fatal(c)
	}
	gap, _ := Parse("silk s 8\nfield a 0 4 TRIPLE P1 bucket\n  code x 4 1800..1900\n")
	if c := gap.Common[0].bucketCodes(edtf("2000")); !slices.Equal(c, []uint64{CodeOther}) {
		t.Fatal(c)
	}
}
