//ff:func feature=cli type=helper control=sequence
//ff:what TestLabelMatches — each language as a LABEL language condition
package cli

import (
	"slices"
	"testing"
)

func TestLabelMatches(t *testing.T) {
	if m := labelMatches([]string{"ko", "en"}); !slices.Equal(m, []string{"LABEL:2.0=@ko", "LABEL:2.0=@en"}) {
		t.Fatal(m)
	}
}
