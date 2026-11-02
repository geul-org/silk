//ff:func feature=sig type=scanner control=sequence
//ff:what TestPutGet — bits land big-endian across bytes; sortRows drops equal rows
package sig

import (
	"bytes"
	"testing"
)

func TestPutGet(t *testing.T) {
	b := make([]byte, 3)
	put(b, 5, 7, 0x55)
	if get(b, 5, 7) != 0x55 || !bytes.Equal(b, []byte{0b00000101, 0b01010000, 0}) {
		t.Fatal(b)
	}
	put(b, 5, 7, 0)
	if b[0]|b[1] != 0 {
		t.Fatal(b)
	}
	if r := sortRows([][]byte{{2}, {1}, {2}}); len(r) != 2 || r[0][0] != 1 {
		t.Fatal(r)
	}
}
