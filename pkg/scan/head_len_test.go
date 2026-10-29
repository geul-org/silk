//ff:func feature=scan type=decoder control=sequence
//ff:what TestHeadLen — 2, 4 or 6 bytes by Length code
package scan

import "testing"

func TestHeadLen(t *testing.T) {
	if headLen([]byte{0x14, 0x8A}) != 2 || headLen([]byte{0x17, 0x0A}) != 4 || headLen([]byte{0x17, 0x8A}) != 6 {
		t.Fatal("headLen")
	}
}
