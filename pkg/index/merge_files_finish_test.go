//ff:func feature=index type=scanner control=sequence
//ff:what TestMergeFilesFinish — an output blocked by a directory fails at rename
package index

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestMergeFilesFinish(t *testing.T) {
	dir := t.TempDir()
	info := seg.Info{RowWidth: 8, DocOff: 4}
	p := filepath.Join(dir, "r")
	w, _ := seg.NewWriter(p, info)
	w.Add(nil, make([]byte, 8))
	w.Finish()
	out := filepath.Join(dir, "o")
	os.MkdirAll(filepath.Join(out, "x"), 0o755)
	if _, err := mergeFiles([]string{p}, info, seg.Hooks{}, out); err == nil {
		t.Fatal("rename over a directory")
	}
	if _, err := mergeKept([]string{p}, info, seg.Hooks{}, out); err == nil {
		t.Fatal("rename over a directory")
	}
}
