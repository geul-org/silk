//ff:func feature=index type=helper control=sequence
//ff:what TestBitmapGet — set bits read true, others and bits beyond N false
package index

import "testing"

func TestBitmapGet(t *testing.T) {
	b := Bitmap{B: []byte{0b101}, N: 3}
	if !b.Get(0) || b.Get(1) || !b.Get(2) || b.Get(3) || b.Get(100) {
		t.Fatal(b)
	}
}
