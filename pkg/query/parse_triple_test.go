//ff:func feature=query type=decoder control=iteration dimension=1
//ff:what TestParseTriple — P106=Q169470 gives canonical records; wildcards allowed on one side; malformed conditions fail
package query

import (
	"bytes"
	"errors"
	"testing"
)

func TestParseTriple(t *testing.T) {
	tr, err := ParseTriple("P106 = Q169470")
	if err != nil || !bytes.Equal(tr.Pred, []byte{0x00, 0xa8, 0x00, 0x6a}) || !bytes.Equal(tr.Obj, []byte{0x01, 0x20, 0x00, 0x02, 0x95, 0xfe}) {
		t.Fatalf("%x %x %v", tr.Pred, tr.Obj, err)
	}
	if tr, err := ParseTriple("P106=*"); err != nil || tr.Obj != nil {
		t.Fatal(tr, err)
	}
	for _, s := range []string{"P106", "X1=Q2", "P1=Qx", "*=*"} {
		if _, err := ParseTriple(s); !errors.Is(err, ErrSyntax) {
			t.Fatal(s, err)
		}
	}
}
