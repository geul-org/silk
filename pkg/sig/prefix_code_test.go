//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what TestPrefixCode — written bits; values optional; bad bits fail
package sig

import "testing"

func TestPrefixCode(t *testing.T) {
	if e, err := prefixCode("n", "0b0101", []string{"Q1"}); err != nil || e.Code != 5 || e.Len != 4 {
		t.Fatal(e, err)
	}
	if e, err := prefixCode("n", "0b01", nil); err != nil || len(e.Values) != 0 {
		t.Fatal(e, err)
	}
	for _, s := range []string{"12", "0b", "0b2"} {
		if _, err := prefixCode("n", s, nil); err == nil {
			t.Fatal(s)
		}
	}
}
