//ff:func feature=sig type=scanner control=sequence
//ff:what TestEnumCode — a number and values; a bad value fails
package sig

import "testing"

func TestEnumCode(t *testing.T) {
	if e, err := enumCode("n", "0x10", []string{"Q1", "Q2"}, 8); err != nil || e.Code != 16 || len(e.Values) != 2 || e.Len != 8 {
		t.Fatal(e, err)
	}
	if _, err := enumCode("n", "4", nil, 8); err == nil {
		t.Fatal("no value")
	}
	if _, err := enumCode("n", "4", []string{"Z"}, 8); err == nil {
		t.Fatal("bad value")
	}
}
