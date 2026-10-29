//ff:func feature=index type=helper control=sequence
//ff:what SpecFile — the path of the spec a manifest was built with
package index

import "path/filepath"

// SpecFile is <dir>/spec/<hash>.json.
func SpecFile(dir string, m Manifest) string {
	return filepath.Join(dir, specName(m.Spec))
}
