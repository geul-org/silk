//ff:func feature=recmap type=writer control=sequence
//ff:what WriteFile — encode a map and put it at path atomically (temporary file, then rename), creating directories
package recmap

import (
	"os"
	"path/filepath"
)

// WriteFile writes m.Encode() to path through path+".tmp".
func WriteFile(path string, m *Map) error {
	b, err := m.Encode()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, b, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}
