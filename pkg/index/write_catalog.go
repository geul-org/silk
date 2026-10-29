//ff:func feature=index type=writer control=sequence
//ff:what writeCatalog — index.json written atomically (temporary file, then rename)
package index

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// writeCatalog puts c at dir/index.json.
func writeCatalog(dir string, c Catalog) error {
	b, _ := json.MarshalIndent(c, "", "  ")
	p := filepath.Join(dir, catalogFile)
	if err := os.WriteFile(p+".tmp", append(b, '\n'), 0o644); err != nil {
		return err
	}
	return os.Rename(p+".tmp", p)
}
