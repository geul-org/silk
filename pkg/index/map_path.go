//ff:func feature=index type=helper control=sequence
//ff:what mapPath — a catalog-relative map file as a path under the index directory
package index

import "path/filepath"

// mapPath joins dir and the slash-separated rel.
func mapPath(dir, rel string) string {
	return filepath.Join(dir, filepath.FromSlash(rel))
}
