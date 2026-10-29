//ff:func feature=index type=reader control=sequence
//ff:what TestReadCatalog — missing is ErrStale; bad JSON and another version are ErrCatalog
package index

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestReadCatalog(t *testing.T) {
	dir := t.TempDir()
	if _, err := readCatalog(dir); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
	p := filepath.Join(dir, catalogFile)
	os.WriteFile(p, []byte("{"), 0o644)
	if _, err := readCatalog(dir); !errors.Is(err, ErrCatalog) {
		t.Fatal(err)
	}
	os.WriteFile(p, []byte(`{"version":9}`), 0o644)
	if _, err := readCatalog(dir); !errors.Is(err, ErrCatalog) {
		t.Fatal(err)
	}
	os.WriteFile(p, []byte(`{"version":1,"maps":["a/b.srm"]}`), 0o644)
	if c, err := readCatalog(dir); err != nil || c.Maps[0] != "a/b.srm" {
		t.Fatal(c, err)
	}
}
