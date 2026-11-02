//ff:func feature=sig type=scanner control=sequence
//ff:what TestEdtfYears — a date, a decade, BCE, junk
package sig

import "testing"

func TestEdtfYears(t *testing.T) {
	if a, b, ok := edtfYears("195X"); !ok || a != 1950 || b != 1959 {
		t.Fatal(a, b, ok)
	}
	if a, b, ok := edtfYears("-0043-03"); !ok || a != -43 || b != -43 {
		t.Fatal(a, b, ok)
	}
	if _, _, ok := edtfYears("abc"); ok {
		t.Fatal("junk")
	}
}
