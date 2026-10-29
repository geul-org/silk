//ff:func feature=spec type=decoder control=sequence
//ff:what TestField — TRIPLE Predicate and Object, a LABEL's language record (short code and BCP 47), and false for a missing field or a sub-record of a Reference
package spec

import (
	"bytes"
	"testing"
)

func TestField(t *testing.T) {
	tr := hexBytes("148a 0088 0005 00a8 001f 00a0 002a 00a0 0005 0000")
	if f, ok := Field(tr, []int{1}); !ok || !bytes.Equal(f, hexBytes("00a8 001f")) {
		t.Fatal(f)
	}
	if f, ok := Field(tr, []int{3}); !ok || !bytes.Equal(f, hexBytes("00a0 0005")) {
		t.Fatal(f)
	}
	if _, ok := Field(tr, []int{5}); ok {
		t.Fatal("field 5")
	}
	if _, ok := Field(tr, []int{1, 0}); ok {
		t.Fatal("sub-record of a Reference")
	}
	lb := hexBytes("1604 0088 0008 00a0 009b 0bca 09c7 7074 2d42 5200 4272 6173 696c")
	if f, ok := Field(lb, []int{2, 0}); !ok || !bytes.Equal(f, hexBytes("09c7 7074 2d42 5200")) {
		t.Fatal(f)
	}
}
