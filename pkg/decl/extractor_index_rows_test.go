//ff:func feature=decl type=scanner control=sequence
//ff:what TestExtractorIndexRows — every record of the TAG, with its number and offset
package decl

import (
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestExtractorIndexRows(t *testing.T) {
	d, _ := NewIndex(fixture.Grammar(t), "t", "TRIPLE", "predicate", []string{"object"}, []int{6})
	b := rowsOfKind(t, d)
	if b[0].Rows() != 3 {
		t.Fatal(b[0].Rows())
	}
}
