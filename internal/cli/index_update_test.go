//ff:func feature=cli type=command control=sequence
//ff:what TestIndexUpdate — the next generation is indexed incrementally; a bad spec file, a non-archive and a missing index fail
package cli

import (
	"io"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/silk/internal/fixture"
)

func TestIndexUpdate(t *testing.T) {
	data := packedIndex(t)
	fixture.Evolve(t, data)
	dir := filepath.Join(data, ".silk")
	var o strings.Builder
	if code := indexUpdate(data, dir, indexFlags{compact: true}, &o, io.Discard); code != 0 || !strings.Contains(o.String(), "new 2, kept 2, removed 2") {
		t.Fatal(code, o.String())
	}
	if code := indexUpdate(data, dir, indexFlags{spec: "/nonexistent"}, io.Discard, io.Discard); code != 1 {
		t.Fatal("spec", code)
	}
	if code := indexUpdate(t.TempDir(), dir, indexFlags{}, io.Discard, io.Discard); code != 1 {
		t.Fatal("archive", code)
	}
	if code := indexUpdate(data, t.TempDir(), indexFlags{}, io.Discard, io.Discard); code != 1 {
		t.Fatal("no index", code)
	}
}
