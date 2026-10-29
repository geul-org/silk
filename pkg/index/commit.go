//ff:func feature=index type=writer control=sequence
//ff:what commit — the next generation: write its manifest, then switch CURRENT to it (the atomic step); bytes written
package index

import (
	"encoding/json"
	"fmt"
	"path/filepath"
)

// commit increments m.Generation and makes m current.
func commit(dir string, m *Manifest) (int64, error) {
	m.Version = manifestVersion
	m.Generation++
	b, _ := json.MarshalIndent(m, "", "  ")
	n1, err := writeAtomic(filepath.Join(dir, manifestName(m.Generation)), append(b, '\n'))
	if err != nil {
		return 0, err
	}
	n2, err := writeAtomic(filepath.Join(dir, currentFile), fmt.Appendf(nil, "%d\n", m.Generation))
	return n1 + n2, err
}
