//ff:func feature=sig type=scanner control=sequence
//ff:what TestParseValues — tokens to records; a bad token stops
package sig

import "testing"

func TestParseValues(t *testing.T) {
	if v, err := parseValues([]string{"Q1", "P2"}); err != nil || len(v) != 2 {
		t.Fatal(v, err)
	}
	if _, err := parseValues([]string{"Q1", "Z"}); err == nil {
		t.Fatal("bad")
	}
}
