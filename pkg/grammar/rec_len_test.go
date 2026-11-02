//ff:func feature=grammar type=scanner control=iteration dimension=1
//ff:what TestRecLen — direct and extended lengths; too short and bad prefix are 0
package grammar

import "testing"

func TestRecLen(t *testing.T) {
	cases := []struct {
		b    []byte
		want int
	}{{[]byte{0x00, 0xA0}, 4}, {[]byte{0x17, 0x0A, 0, 20}, 40}, {[]byte{0x17, 0x8A, 0, 1, 0, 0}, 1 << 17}, {[]byte{0x17, 0x0A}, 0}, {[]byte{0x17, 0x8A, 0, 0}, 0}, {[]byte{0}, 0}, {[]byte{0xFF, 0}, 0}}
	for _, c := range cases {
		if n := RecLen(c.b); n != c.want {
			t.Fatal(c.b, n)
		}
	}
	if headLen([]byte{0x17, 0x0A}) != 4 || headLen([]byte{0x17, 0x8A}) != 6 || headLen([]byte{0x00, 0xA0}) != 2 {
		t.Fatal("headLen")
	}
}
