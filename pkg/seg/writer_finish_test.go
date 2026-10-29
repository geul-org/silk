//ff:func feature=seg type=writer control=sequence
//ff:what TestWriterFinish — the finished file opens with the same Info and size; a failing file is removed and reported
package seg

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriterFinish(t *testing.T) {
	p := filepath.Join(t.TempDir(), "a.sseg")
	w, _ := NewWriter(p, Info{Tag: 7, PartWidth: 2, SortWidth: 2, Sorted: true}, 1)
	w.Begin([]byte("aa"))
	w.Add(row("k1", 1, 0))
	w.End()
	info, err := w.Finish()
	f, err2 := Open(p)
	if err != nil || err2 != nil || f.Info != info || info.Bytes != w.Written() || info.Parts != 1 {
		t.Fatal(err, err2, info)
	}
	f.Close()
	b := brokenWriter(t, 1)
	if _, err := b.Finish(); err == nil {
		t.Fatal("closed file")
	}
	if _, err := os.Stat(b.Path + ".tmp"); !os.IsNotExist(err) {
		t.Fatal("temporary file kept")
	}
}
