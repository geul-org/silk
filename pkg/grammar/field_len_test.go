//ff:func feature=grammar type=scanner control=sequence
//ff:what TestFieldLen — every width form, and fields past the record
package grammar

import "testing"

func TestFieldLen(t *testing.T) {
	r := []byte{0, 0xA0, 0, 1, 0, 2}
	if n := fieldLen(r, Field{}, 0, 0); n != 4 {
		t.Fatal(n)
	}
	if n := fieldLen(r, Field{Width: 1}, 4, 0); n != 2 {
		t.Fatal(n)
	}
	if n := fieldLen(r, Field{Width: RestWidth}, 2, 2); n != 2 {
		t.Fatal(n)
	}
	if n := fieldLen(r, Field{}, 9, 0); n != -1 {
		t.Fatal(n)
	}
	if n := fieldLen([]byte{0xFF, 0xFF}, Field{}, 0, 0); n != -1 {
		t.Fatal(n)
	}
	if n := fieldLen(r, Field{Width: RestWidth}, 2, 6); n != -1 {
		t.Fatal(n)
	}
	if n := fieldLen(r, Field{Width: 9}, 0, 0); n != -1 {
		t.Fatal(n)
	}
}
