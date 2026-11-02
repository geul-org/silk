//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what TestIngesterSinkFlush — past the budget a failing flush surfaces
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIngesterSinkFlush(t *testing.T) {
	_, dir, a := built(t, 1)
	u := &updater{ar: a, dir: dir, o: testOpts, m: current(t, dir), tmp: t.TempDir()}
	ts, _ := u.targets()
	g := &ingester{u: u, ts: ts, budget: 1}
	loc := g.locals()
	for _, b := range loc[0] {
		b.Add(nil, make([]byte, ts[0].sp[0].info.RowWidth))
	}
	file := filepath.Join(dir, "f")
	os.WriteFile(file, nil, 0o644)
	ts[0].sp[0].tmp = file
	if err := g.sink(loc, 1); err == nil {
		t.Fatal("flush under a file")
	}
}
