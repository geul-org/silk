//ff:func feature=cli type=command control=sequence
//ff:what TestIndexVerify — a built index verifies; a non-archive or a missing index fails
package cli

import (
	"io"
	"path/filepath"
	"strings"
	"testing"
)

func TestIndexVerify(t *testing.T) {
	data := packedIndex(t)
	var o strings.Builder
	if code := indexVerify(data, filepath.Join(data, ".silk"), &o, io.Discard); code != 0 || !strings.Contains(o.String(), "4 documents") {
		t.Fatal(code, o.String())
	}
	if code := indexVerify(t.TempDir(), filepath.Join(data, ".silk"), io.Discard, io.Discard); code != 1 {
		t.Fatal(code)
	}
	if code := indexVerify(data, t.TempDir(), io.Discard, io.Discard); code != 1 {
		t.Fatal(code)
	}
}
