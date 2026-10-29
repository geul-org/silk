//ff:func feature=spec type=decoder control=iteration dimension=1
//ff:what TestRecLen — direct lengths, 1- and 2-word extensions, too-short input and a non-000 prefix
package spec

import "testing"

func TestRecLen(t *testing.T) {
	cases := []struct {
		b    string
		want int
	}{
		{"14", 0},
		{"148a", 20},
		{"2000", -1},
		{"170a 0026", 76},
		{"170a", 0},
		{"178a 0001 0000", 0x20000},
		{"178a 0001", 0},
	}
	for _, c := range cases {
		if got := RecLen(hexBytes(c.b)); got != c.want {
			t.Fatalf("%s: %d", c.b, got)
		}
	}
}
