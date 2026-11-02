//ff:func feature=sig type=scanner control=sequence
//ff:what TestDefExactTerm — the when code alone, exact; the candidate term holds somevalue and other
package sig

import "testing"

func TestDefExactTerm(t *testing.T) {
	d := testDef(t)
	x := d.exactTerm(d.Whens[0])
	if x.Spans[0].Lo != 4 || x.Exact[0].Hi != 4 {
		t.Fatal(x)
	}
	if y := d.whenTerm(d.Common[0], false); y.Spans[0] != (Span{CodeSomeValue, CodeOther}) || y.Exact != nil {
		t.Fatal(y)
	}
}
