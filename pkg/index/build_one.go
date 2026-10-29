//ff:func feature=index type=builder control=sequence
//ff:what buildOne — build one target's map from its data pack file, name its Issuer, generation, manifest and pack, and write it
package index

import (
	"os"

	"github.com/park-jun-woo/silk/pkg/recmap"
)

// buildOne writes t's map to path.
func buildOne(path string, t Target) (recmap.Header, error) {
	f, err := os.Open(t.Pack.Path)
	if err != nil {
		return recmap.Header{}, err
	}
	defer f.Close()
	m, err := recmap.Build(f, t.Pack.Size, t.Docs)
	if err != nil {
		return recmap.Header{}, err
	}
	m.Header.Issuer, m.Header.Generation = t.Issuer, t.Pack.Set.Generation
	m.Header.Manifest, m.Header.Pack = t.Pack.Set.Manifest, t.Pack.Hash
	if err := recmap.WriteFile(path, m); err != nil {
		return recmap.Header{}, err
	}
	return m.Header, nil
}
