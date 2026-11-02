//ff:func feature=seg type=scanner control=sequence
//ff:what TestWriterFinish — table failure, copy failure and sync failure abort; an empty segment is valid
package seg

import (
	"bufio"
	"os"
	"path/filepath"
	"testing"
)

func TestWriterFinish(t *testing.T) {
	p := filepath.Join(t.TempDir(), "empty")
	w, _ := NewWriter(p, testInfo())
	if n, err := w.Finish(); n != 0 || err != nil {
		t.Fatal(n, err)
	}
	f, err := Open(p)
	if err != nil || f.Rows != 0 || f.Parts != 0 || f.Verify() != nil {
		t.Fatal(err)
	}
	f.Close()
	dir := t.TempDir()
	w, _ = NewWriter(filepath.Join(dir, "a"), testInfo())
	w.tw = bufio.NewWriterSize(failWriter{}, 8)
	w.Add([]byte("pa"), row(0, 1, 1))
	if _, err := w.Finish(); err == nil {
		t.Fatal("table failure")
	}
	w, _ = NewWriter(filepath.Join(dir, "b"), testInfo())
	w.tf.Close()
	if _, err := w.Finish(); err == nil {
		t.Fatal("copy failure")
	}
	w, _ = NewWriter(filepath.Join(dir, "c"), testInfo())
	w.f.Close()
	if _, err := w.Finish(); err == nil {
		t.Fatal("close failure")
	}
	if _, err := os.Stat(filepath.Join(dir, "c.tmp")); err == nil {
		t.Fatal("temporary file left")
	}
}
