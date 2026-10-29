//ff:func feature=archive type=helper control=sequence
//ff:what wellKnown — the `.well-known/geul/` directory of a data folder, or the directory itself when it has none
package archive

import (
	"os"
	"path/filepath"
)

// wellKnown returns dir/.well-known/geul when it is a directory, else dir.
func wellKnown(dir string) string {
	wk := filepath.Join(dir, ".well-known", "geul")
	if fi, err := os.Stat(wk); err == nil && fi.IsDir() {
		return wk
	}
	return dir
}
