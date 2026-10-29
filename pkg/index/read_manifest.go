//ff:func feature=index type=reader control=sequence
//ff:what readManifest — one generation's manifest, format version checked
package index

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// readManifest reads manifest/<gen>.json under dir.
func readManifest(dir string, gen uint64) (Manifest, error) {
	b, err := os.ReadFile(filepath.Join(dir, manifestName(gen)))
	if err != nil {
		return Manifest{}, err
	}
	var m Manifest
	if err := json.Unmarshal(b, &m); err != nil || m.Version != manifestVersion || m.Generation != gen {
		return Manifest{}, fmt.Errorf("%w: manifest %d", ErrFormat, gen)
	}
	return m, nil
}
