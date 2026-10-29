//ff:func feature=index type=reader control=sequence
//ff:what TestPinnedGens — generations of live pins; malformed pins and pins of dead processes are removed
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPinnedGens(t *testing.T) {
	d := t.TempDir()
	pinGen(d, 3)
	pd := filepath.Join(d, dirPin)
	os.WriteFile(filepath.Join(pd, "junk"), nil, 0o644)
	os.WriteFile(filepath.Join(pd, "5-1073741824-1"), nil, 0o644)
	g := pinnedGens(d)
	es, _ := os.ReadDir(pd)
	if len(g) != 1 || !g[3] || len(es) != 1 {
		t.Fatal(g, len(es))
	}
}
