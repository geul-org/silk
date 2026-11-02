//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what TestFieldReservedTerm — codes keep other as a candidate where a limited row can hide them
package sig

import "testing"

func TestFieldReservedTerm(t *testing.T) {
	f := &Field{}
	if x := f.reservedTerm(CodeNone); len(x.Spans) != 1 || len(x.Exact) != 1 {
		t.Fatal(x)
	}
	if x := f.reservedTerm(CodeOther); x.Exact != nil {
		t.Fatal(x)
	}
	for _, c := range []uint64{CodeNoValue, CodeSomeValue} {
		if x := f.reservedTerm(c); len(x.Spans) != 2 || x.Spans[1].Lo != CodeOther {
			t.Fatal(x)
		}
	}
}
