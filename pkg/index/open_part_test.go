//ff:func feature=index type=reader control=sequence
//ff:what TestOpenPart — a map pairs with its pack; missing map, unknown manifest and missing pack fail
package index

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/silk/pkg/archive"
)

func TestOpenPart(t *testing.T) {
	a, dir := fixtureArchive(t)
	Build(context.Background(), a, dir)
	c, _ := readCatalog(dir)
	p := mapPath(dir, c.Maps[0])
	part, err := openPart(a, p)
	if err != nil || part.Map.Header.Docs != 4 {
		t.Fatal(err)
	}
	part.Map.Close()
	part.Data.Close()
	if _, err := openPart(a, filepath.Join(dir, "none.srm")); err == nil {
		t.Fatal("missing map")
	}
	if _, err := openPart(&archive.Archive{}, p); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
	s := a.Sets[0]
	s.Dir = t.TempDir()
	os.MkdirAll(filepath.Join(s.Dir, "o"), 0o755)
	if _, err := openPart(&archive.Archive{Sets: []archive.Set{s}}, p); err == nil {
		t.Fatal("missing pack")
	}
}
