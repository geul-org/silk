//ff:func feature=grammar type=helper control=sequence
//ff:what repoRoot — the module root
package grammar

import (
	"path/filepath"
	"runtime"
)

func repoRoot() string {
	_, f, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(f), "..", "..")
}
