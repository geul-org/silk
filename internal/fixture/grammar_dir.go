//ff:func feature=fixture type=helper control=sequence
//ff:what GrammarDir — the copy of the published `0x0000` pack set
package fixture

import (
	"path/filepath"
	"runtime"
)

func GrammarDir() string {
	_, f, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(f), "..", "..", "testdata", "grammar", "0x0000")
}
