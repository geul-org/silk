// Package fixture builds SILK's test archives from testdata/archive.
//
//ff:func feature=fixture type=helper control=sequence
//ff:what Source — the testdata/archive data folder (4 Wikidata version documents under Issuer 0x00010000, 2 with P106=Q169470)
package fixture

import (
	"path/filepath"
	"runtime"
)

// Source locates testdata/archive from this file's path.
func Source() string {
	_, f, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(f), "..", "..", "testdata", "archive")
}
