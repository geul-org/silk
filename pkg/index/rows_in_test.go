//ff:func feature=index type=reader control=sequence
//ff:what TestRowsIn — a segment's rows; a missing file fails
package index

import (
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestRowsIn(t *testing.T) {
	p := filepath.Join(t.TempDir(), "r")
	w, _ := seg.NewWriter(p, pathInfo())
	w.Add(nil, make([]byte, pathKey+4))
	w.Finish()
	if n, err := rowsIn(p); n != 1 || err != nil {
		t.Fatal(n, err)
	}
	if _, err := rowsIn(p + "x"); err == nil {
		t.Fatal("missing")
	}
}
