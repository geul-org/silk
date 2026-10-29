//ff:func feature=index type=helper control=sequence
//ff:what TestBitmapCount — the number of set bits
package index

import "testing"

func TestBitmapCount(t *testing.T) {
	if n := (Bitmap{B: []byte{0xF0, 0x01}, N: 9}).Count(); n != 5 {
		t.Fatal(n)
	}
}
