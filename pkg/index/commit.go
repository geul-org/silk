//ff:func feature=index type=writer control=sequence
//ff:what commit — the next generation: its manifest, then CURRENT (the atomic step)
package index

import (
	"encoding/json"
	"fmt"
	"path/filepath"
)

func commit(dir string, m *Manifest) error {
	m.Version = manifestVersion
	m.Generation++
	b, _ := json.MarshalIndent(m, "", "  ")
	if err := writeAtomic(filepath.Join(dir, manifestName(m.Generation)), append(b, '\n')); err != nil {
		return err
	}
	return writeAtomic(filepath.Join(dir, currentFile), fmt.Appendf(nil, "%d\n", m.Generation))
}
