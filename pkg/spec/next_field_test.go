//ff:func feature=spec type=decoder control=sequence
//ff:what TestNextField — the end of a framed field; false at position 0, at the end, for a bad TAG and for a field running past the record
package spec

import "testing"

func TestNextField(t *testing.T) {
	rec := hexBytes("148a 0088 0005 00a8 001f 00a0 002a 00a0 0005 0000")
	if e, ok := nextField(rec, 2); !ok || e != 6 {
		t.Fatal(e, ok)
	}
	bad := hexBytes("148a ffff 0000")
	long := hexBytes("148a 0108 0001")
	if _, ok := nextField(rec, 0); ok {
		t.Fatal("p=0")
	}
	if _, ok := nextField(rec, len(rec)); ok {
		t.Fatal("end")
	}
	if _, ok := nextField(bad, 2); ok {
		t.Fatal("bad prefix")
	}
	if _, ok := nextField(long, 2); ok {
		t.Fatal("past the end")
	}
}
