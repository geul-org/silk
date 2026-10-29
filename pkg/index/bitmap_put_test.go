//ff:func feature=index type=helper control=sequence
//ff:what TestBitmapPut — setting beyond N grows the bitmap; clearing works
package index

import "testing"

func TestBitmapPut(t *testing.T) {
	var b Bitmap
	b.Put(9, true)
	b.Put(2, true)
	b.Put(2, false)
	if b.N != 10 || len(b.B) != 2 || !b.Get(9) || b.Get(2) {
		t.Fatal(b)
	}
}
