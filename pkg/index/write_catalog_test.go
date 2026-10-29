//ff:func feature=index type=writer control=sequence
//ff:what TestWriteCatalog — written catalog reads back; temporary-file and rename failures surface
package index

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestWriteCatalog(t *testing.T) {
	dir := t.TempDir()
	c := Catalog{Version: catalogVersion, Sets: []CatalogSet{{Name: "a"}}, Maps: []string{"a/b.srm"}}
	if err := writeCatalog(dir, c); err != nil {
		t.Fatal(err)
	}
	if got, err := readCatalog(dir); err != nil || !slices.Equal(got.Maps, c.Maps) || got.Sets[0] != c.Sets[0] {
		t.Fatal(got, err)
	}
	if err := writeCatalog(filepath.Join(dir, "none"), c); err == nil {
		t.Fatal("missing directory")
	}
	d2 := t.TempDir()
	os.MkdirAll(filepath.Join(d2, catalogFile, "x"), 0o755)
	if err := writeCatalog(d2, c); err == nil {
		t.Fatal("rename onto a non-empty directory")
	}
}
