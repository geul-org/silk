//ff:func feature=index type=decoder control=sequence
//ff:what TestRecLenAt — the length of the record at a byte position
package index

import "testing"

func TestRecLenAt(t *testing.T) {
	if n := recLenAt([]byte{0, 0, 0x14, 0x8a}, 2); n != 20 {
		t.Fatal(n)
	}
}
