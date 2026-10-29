//ff:func feature=seg type=writer control=sequence
//ff:what TestWriterAbort — the temporary file is gone and no segment appears
package seg

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriterAbort(t *testing.T) {
	p := filepath.Join(t.TempDir(), "a.sseg")
	w, _ := NewWriter(p, Info{}, 1)
	w.Abort()
	_, e1 := os.Stat(p + ".tmp")
	_, e2 := os.Stat(p)
	if !os.IsNotExist(e1) || !os.IsNotExist(e2) {
		t.Fatal(e1, e2)
	}
}
