//ff:func feature=index type=helper control=sequence
//ff:what TestRowsOf — the sum of rows
package index

import "testing"

func TestRowsOf(t *testing.T) {
	if rowsOf([]SegRef{{Rows: 2}, {Rows: 5}}) != 7 || rowsOf(nil) != 0 {
		t.Fatal("rowsOf")
	}
}
