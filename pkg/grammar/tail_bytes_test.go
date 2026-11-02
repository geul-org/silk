//ff:func feature=grammar type=scanner control=sequence
//ff:what TestTailBytes — fixed words after a rest field count; fieldLen refuses what does not fit
package grammar

import "testing"

func TestTailBytes(t *testing.T) {
	if n := tailBytes([]Field{{Width: 1}, {Width: RestWidth}, {Width: 2}}); n != 4 {
		t.Fatal(n)
	}
	if fieldLen([]byte{0, 0, 0, 0}, Field{Width: 1}, 2, 0) != 2 {
		t.Fatal("raw")
	}
	if fieldLen([]byte{0xFF, 0xFF}, Field{}, 0, 0) != -1 {
		t.Fatal("bad record")
	}
	if fieldLen([]byte{0, 0}, Field{Width: 9}, 0, 0) != -1 {
		t.Fatal("wide")
	}
	if fieldLen([]byte{0, 0}, Field{Width: RestWidth}, 2, 4) != -1 {
		t.Fatal("tail")
	}
}
