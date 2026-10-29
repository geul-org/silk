//ff:func feature=cli type=helper control=sequence
//ff:what TestSegBytes — bytes of runs and appended segments
package cli

import (
	"testing"

	"github.com/geul-org/silk/pkg/index"
)

func TestSegBytes(t *testing.T) {
	if n := segBytes(index.Section{Runs: []index.SegRef{{Bytes: 3}}, Appends: []index.SegRef{{Bytes: 4}}}); n != 7 {
		t.Fatal(n)
	}
}
