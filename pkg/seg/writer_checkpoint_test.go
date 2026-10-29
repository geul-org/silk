//ff:func feature=seg type=writer control=sequence
//ff:what TestWriterCheckpoint — between partitions a State with bytes, table and hash; inside a partition or on a failing file an error
package seg

import (
	"errors"
	"path/filepath"
	"testing"
)

func TestWriterCheckpoint(t *testing.T) {
	w, _ := NewWriter(filepath.Join(t.TempDir(), "a.sseg"), Info{PartWidth: 2, SortWidth: 2}, 1)
	defer w.Abort()
	w.Begin([]byte("aa"))
	if _, err := w.Checkpoint(); !errors.Is(err, ErrOrder) {
		t.Fatal(err)
	}
	w.Add(row("k1", 1, 0))
	w.End()
	st, err := w.Checkpoint()
	if err != nil || st.Written != 10 || len(st.Parts) != 1 || len(st.Hash) == 0 || string(st.Key) != "aa" {
		t.Fatal(err, st)
	}
	b := brokenWriter(t, 1)
	b.write([]byte{1})
	if _, err := b.Checkpoint(); err == nil {
		t.Fatal("closed file")
	}
}
