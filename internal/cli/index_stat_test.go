//ff:func feature=cli type=command control=sequence
//ff:what TestIndexStat — sections with runs, appended segments, rows and bytes, and due merges; a missing index fails
package cli

import (
	"io"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/silk/internal/fixture"
)

func TestIndexStat(t *testing.T) {
	data := packedIndex(t)
	fixture.Evolve(t, data)
	call("index update " + data)
	var o strings.Builder
	if code := indexStat(filepath.Join(data, ".silk"), &o, io.Discard); code != 0 || !strings.Contains(o.String(), "appended 1") || !strings.Contains(o.String(), "due: TRIPLE") {
		t.Fatal(code, o.String())
	}
	if code := indexStat(t.TempDir(), io.Discard, io.Discard); code != 1 {
		t.Fatal(code)
	}
}
