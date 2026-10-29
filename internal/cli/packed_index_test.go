//ff:func feature=cli type=helper control=sequence
//ff:what packedIndex — test helper: the packed test archive with its index built by `silk index build` (minimum partition from the default spec)
package cli

import (
	"testing"

	"github.com/park-jun-woo/silk/internal/fixture"
)

// packedIndex returns the data folder; the index is at <data>/.silk.
func packedIndex(t testing.TB) string {
	t.Helper()
	data := fixture.Packed(t, nil)
	if code, _, e := call("index build " + data); code != 0 {
		t.Fatal(e)
	}
	return data
}
