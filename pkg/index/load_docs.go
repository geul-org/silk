//ff:func feature=index type=reader control=iteration dimension=1
//ff:what loadDocs — every document table file of a manifest in order; each must continue the sequence numbers
package index

import (
	"fmt"
	"path/filepath"
)

// loadDocs concatenates the document tables.
func loadDocs(dir string, m Manifest) ([]DocEntry, error) {
	var out []DocEntry
	for _, f := range m.Docs {
		first, ds, err := readDocs(filepath.Join(dir, f.File))
		if err != nil {
			return nil, err
		}
		if first != uint32(len(out)) || first != f.First || uint32(len(ds)) != f.Count {
			return nil, fmt.Errorf("%w: %s holds %d from %d", ErrFormat, f.File, len(ds), first)
		}
		out = append(out, ds...)
	}
	return out, nil
}
