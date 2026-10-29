//ff:func feature=seg type=writer control=iteration dimension=1
//ff:what TestWriterBegin — keys must ascend, have the partition width, and not open inside another partition
package seg

import (
	"errors"
	"path/filepath"
	"testing"
)

func TestWriterBegin(t *testing.T) {
	w, _ := NewWriter(filepath.Join(t.TempDir(), "a.sseg"), Info{PartWidth: 2}, 1)
	defer w.Abort()
	if err := w.Begin([]byte("bb")); err != nil {
		t.Fatal(err)
	}
	if err := w.Begin([]byte("cc")); !errors.Is(err, ErrOrder) {
		t.Fatal("inside a partition", err)
	}
	w.End()
	for _, k := range []string{"b", "bb", "aa"} {
		if err := w.Begin([]byte(k)); !errors.Is(err, ErrOrder) {
			t.Fatal(k, err)
		}
	}
	if err := w.Begin([]byte("cc")); err != nil {
		t.Fatal(err)
	}
}
