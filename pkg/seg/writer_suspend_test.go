//ff:func feature=seg type=writer control=sequence
//ff:what TestWriterSuspend — a State and a closed file; inside a partition an error
package seg

import (
	"path/filepath"
	"testing"
)

func TestWriterSuspend(t *testing.T) {
	w, _ := NewWriter(filepath.Join(t.TempDir(), "a.sseg"), Info{PartWidth: 2}, 1)
	if st, err := w.Suspend(); err != nil || st.Path == "" {
		t.Fatal(err)
	}
	if err := w.f.Close(); err == nil {
		t.Fatal("file still open")
	}
	w2, _ := NewWriter(filepath.Join(t.TempDir(), "b.sseg"), Info{PartWidth: 2}, 1)
	w2.Begin([]byte("aa"))
	if _, err := w2.Suspend(); err == nil {
		t.Fatal("inside a partition")
	}
}
