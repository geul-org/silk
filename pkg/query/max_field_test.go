//ff:func feature=query type=helper control=sequence
//ff:what TestMaxField — the highest field number, 0 without fields
package query

import "testing"

func TestMaxField(t *testing.T) {
	if maxField(Cond{Fields: []FieldCond{{Path: []int{3}}, {Path: []int{1}}}}) != 3 || maxField(Cond{}) != 0 {
		t.Fatal("maxField")
	}
}
