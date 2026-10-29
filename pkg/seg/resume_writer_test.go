//ff:func feature=seg type=writer control=sequence
//ff:what TestResumeWriter — a resumed writer (bytes past the checkpoint dropped) finishes the same file as an uninterrupted one; a missing or short file or a bad hash state fails
package seg

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestResumeWriter(t *testing.T) {
	d := t.TempDir()
	write := func(w *Writer, keys ...string) {
		for _, k := range keys {
			w.Begin([]byte(k))
			w.Add(row("k1", 1, 0))
			w.End()
		}
	}
	info := Info{PartWidth: 2, SortWidth: 2, Sorted: true}
	a, _ := NewWriter(filepath.Join(d, "a.sseg"), info, 1)
	write(a, "aa", "bb", "cc")
	a.Finish()
	b, _ := NewWriter(filepath.Join(d, "b.sseg"), info, 1)
	write(b, "aa")
	st, _ := b.Suspend()
	tf, _ := os.OpenFile(filepath.Join(d, "b.sseg.tmp"), os.O_WRONLY|os.O_APPEND, 0)
	tf.Write([]byte{9, 9})
	tf.Close()
	r, err := ResumeWriter(st)
	if err != nil {
		t.Fatal(err)
	}
	write(r, "bb", "cc")
	r.Finish()
	x, _ := os.ReadFile(filepath.Join(d, "a.sseg"))
	y, _ := os.ReadFile(filepath.Join(d, "b.sseg"))
	if !bytes.Equal(x, y) {
		t.Fatal("resumed file differs")
	}
	if _, err := ResumeWriter(State{Path: filepath.Join(d, "none")}); err == nil {
		t.Fatal("missing file")
	}
	os.WriteFile(filepath.Join(d, "c.sseg.tmp"), []byte{1}, 0o644)
	if _, err := ResumeWriter(State{Path: filepath.Join(d, "c.sseg"), Written: 5}); err == nil {
		t.Fatal("short file")
	}
	if _, err := ResumeWriter(State{Path: filepath.Join(d, "c.sseg"), Hash: []byte("bad")}); err == nil {
		t.Fatal("bad hash state")
	}
}
