//ff:func feature=recmap type=helper control=sequence
//ff:what sampleMap — the fixture pack's map with its identity filled in, as index.Build would make it
package recmap

import (
	"os"
	"testing"

	"github.com/park-jun-woo/silk/pkg/archive"
)

// sampleMap builds the map of the fixture pack.
func sampleMap(t *testing.T) (*Map, archive.Pack, *os.File) {
	t.Helper()
	p, f := fixturePack(t)
	m, err := Build(f, p.Size, p.Docs)
	if err != nil {
		t.Fatal(err)
	}
	m.Header.Issuer, m.Header.Generation = "0x00010000", 1
	m.Header.Manifest, m.Header.Pack = p.Set.Manifest, p.Hash
	return m, p, f
}
