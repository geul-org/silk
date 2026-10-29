//ff:func feature=index type=helper control=sequence
//ff:what block — test helper: replace dir/sub with a plain file so creating anything under it fails
package index

import (
	"os"
	"path/filepath"
)

// block makes dir/sub unusable as a directory.
func block(dir, sub string) {
	os.RemoveAll(filepath.Join(dir, sub))
	os.MkdirAll(dir, 0o755)
	os.WriteFile(filepath.Join(dir, sub), nil, 0o644)
}
