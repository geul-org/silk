//ff:func feature=index type=rule control=sequence
//ff:what verifyOne — one map: header names the target's Issuer, generation, manifest and pack; tables match the pack (recmap File.Verify)
package index

import (
	"fmt"
	"os"

	"github.com/park-jun-woo/silk/pkg/recmap"
)

// verifyOne opens the map at path and checks it against t.
func verifyOne(path string, t Target) error {
	m, err := recmap.Open(path)
	if err != nil {
		return err
	}
	defer m.Close()
	h := m.Header
	if h.Issuer != t.Issuer || h.Generation != t.Pack.Set.Generation || h.Manifest != t.Pack.Set.Manifest || h.Pack != t.Pack.Hash {
		return fmt.Errorf("%w: %s", ErrStale, path)
	}
	f, err := os.Open(t.Pack.Path)
	if err != nil {
		return err
	}
	defer f.Close()
	return m.Verify(f, t.Docs)
}
