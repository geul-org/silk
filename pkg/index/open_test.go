//ff:func feature=index type=reader control=sequence
//ff:what TestOpen — a built index opens pinned at its generation; no index, or an archive that moved on, fails and leaves no pin
package index

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestOpen(t *testing.T) {
	a, data, dir := built(t)
	ix, err := Open(a, dir)
	if err != nil || ix.M.Generation == 0 || len(ix.Docs) != 4 || ix.Live.Count() != 4 || len(ix.Segs) != 7 {
		t.Fatal(err)
	}
	ix.Close()
	if _, err := Open(a, t.TempDir()); !errors.Is(err, ErrNoIndex) {
		t.Fatal(err)
	}
	b := evolved(t, data)
	if _, err := Open(b, dir); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
	if es, _ := os.ReadDir(filepath.Join(dir, dirPin)); len(es) != 0 {
		t.Fatal("pin left behind")
	}
}
