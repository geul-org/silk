//ff:func feature=index type=writer control=iteration dimension=1
//ff:what removeStale — delete map files of the index directory that the catalog no longer lists
package index

import (
	"os"
	"path/filepath"
	"slices"
)

// removeStale removes `<dir>/*/*.srm` files not in keep (relative, slash-separated).
func removeStale(dir string, keep []string) error {
	paths, _ := filepath.Glob(filepath.Join(dir, "*", "*"+mapExt))
	for _, p := range paths {
		rel, _ := filepath.Rel(dir, p)
		if slices.Contains(keep, filepath.ToSlash(rel)) {
			continue
		}
		if err := os.Remove(p); err != nil {
			return err
		}
	}
	return nil
}
