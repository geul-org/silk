//ff:func feature=cli type=helper control=sequence
//ff:what TestSegRows — rows of runs and appended segments
package cli

import (
	"testing"

	"github.com/park-jun-woo/silk/pkg/index"
)

func TestSegRows(t *testing.T) {
	if n := segRows(index.Section{Runs: []index.SegRef{{Rows: 3}}, Appends: []index.SegRef{{Rows: 4}}}); n != 7 {
		t.Fatal(n)
	}
}
