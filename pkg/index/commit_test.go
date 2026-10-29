//ff:func feature=index type=writer control=sequence
//ff:what TestCommit — the generation advances and CURRENT follows; a blocked manifest directory or CURRENT path fails
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCommit(t *testing.T) {
	d := t.TempDir()
	m := Manifest{Generation: 4}
	if n, err := commit(d, &m); err != nil || n == 0 || m.Generation != 5 || m.Version != manifestVersion {
		t.Fatal(err, m)
	}
	if g, _ := readCurrent(d); g != 5 {
		t.Fatal(g)
	}
	block(d, dirManifest)
	if _, err := commit(d, &m); err == nil {
		t.Fatal("blocked manifest directory")
	}
	e := t.TempDir()
	os.Mkdir(filepath.Join(e, currentFile+".tmp"), 0o755)
	if _, err := commit(e, &m); err == nil {
		t.Fatal("blocked CURRENT")
	}
}
