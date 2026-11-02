//ff:func feature=decl type=scanner control=sequence
//ff:what TestExtractorSilkRows — both bundles get each row
package decl

import "testing"

func TestExtractorSilkRows(t *testing.T) {
	d, _ := NewSilk("s", silkText)
	b := rowsOfKind(t, d)
	if b[0].Rows() != 2 || b[1].Rows() != 2 {
		t.Fatal(b[0].Rows(), b[1].Rows())
	}
}
