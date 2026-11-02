//ff:func feature=index type=reader control=sequence
//ff:what readManifest — a generation's manifest
package index

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func readManifest(dir string, gen uint64) (Manifest, error) {
	b, err := os.ReadFile(filepath.Join(dir, manifestName(gen)))
	if err != nil {
		return Manifest{}, err
	}
	var m Manifest
	if err := json.Unmarshal(b, &m); err != nil {
		return Manifest{}, fmt.Errorf("%w: manifest: %v", ErrFormat, err)
	}
	if m.Version != manifestVersion {
		return Manifest{}, fmt.Errorf("%w: manifest version %d", ErrFormat, m.Version)
	}
	return m, nil
}
