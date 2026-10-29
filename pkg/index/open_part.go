//ff:func feature=index type=reader control=sequence
//ff:what openPart — one map file and its data pack, the pack found in the set whose manifest the map names
package index

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/silk/pkg/archive"
	"github.com/park-jun-woo/silk/pkg/recmap"
)

// openPart opens the map at path; the set it names must be current in a.
func openPart(a *archive.Archive, path string) (Part, error) {
	m, err := recmap.Open(path)
	if err != nil {
		return Part{}, err
	}
	s, ok := setOf(a, m.Header.Manifest)
	if !ok {
		m.Close()
		return Part{}, fmt.Errorf("%w: %s names manifest %s", ErrStale, path, m.Header.Manifest)
	}
	f, err := os.Open(filepath.Join(s.Dir, "o", m.Header.Pack+".gpk"))
	if err != nil {
		m.Close()
		return Part{}, err
	}
	return Part{Map: m, Data: f}, nil
}
