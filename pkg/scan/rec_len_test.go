//ff:func feature=scan type=decoder control=iteration dimension=1
//ff:what TestRecLen — direct, 1-word and 2-word Length; too short to tell is 0, a non-000 prefix is -1
package scan

import "testing"

func TestRecLen(t *testing.T) {
	cases := map[string]struct {
		b    []byte
		want int
	}{
		"direct":    {[]byte{0x14, 0x8A}, 20},
		"ext1":      {[]byte{0x17, 0x0A, 0x00, 0x20}, 64},
		"ext2":      {[]byte{0x17, 0x8A, 0x00, 0x01, 0x00, 0x00}, 0x20000},
		"short":     {[]byte{0x14}, 0},
		"ext1short": {[]byte{0x17, 0x0A, 0x00}, 0},
		"ext2short": {[]byte{0x17, 0x8A, 0x00, 0x01}, 0},
		"prefix":    {[]byte{0x20, 0x00}, -1},
	}
	for name, c := range cases {
		if got := RecLen(c.b); got != c.want {
			t.Fatal(name, got)
		}
	}
}
