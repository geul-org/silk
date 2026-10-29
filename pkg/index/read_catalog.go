//ff:func feature=index type=reader control=sequence
//ff:what readCatalog — dir/index.json; a missing one says to build the index, a bad one is ErrCatalog
package index

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// readCatalog reads and checks the catalog's format version.
func readCatalog(dir string) (Catalog, error) {
	b, err := os.ReadFile(filepath.Join(dir, catalogFile))
	if err != nil {
		return Catalog{}, fmt.Errorf("%w: %v", ErrStale, err)
	}
	var c Catalog
	if err := json.Unmarshal(b, &c); err != nil || c.Version != catalogVersion {
		return Catalog{}, fmt.Errorf("%w: %s", ErrCatalog, filepath.Join(dir, catalogFile))
	}
	return c, nil
}
