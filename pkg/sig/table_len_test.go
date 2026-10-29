//ff:func feature=sig type=helper control=sequence
//ff:what TestTableLen — rows counted by width; a zero-width table has none
package sig

import "testing"

func TestTableLen(t *testing.T) {
	if (&Table{Words: 2, Data: make([]uint64, 6)}).Len() != 3 || (&Table{}).Len() != 0 {
		t.Fatal("len")
	}
}
