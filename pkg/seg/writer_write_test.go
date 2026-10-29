//ff:func feature=seg type=writer control=sequence
//ff:what TestWriterWrite — bytes counted and hashed; a failing file reports the error
package seg

import (
	"path/filepath"
	"testing"
)

func TestWriterWrite(t *testing.T) {
	w, _ := NewWriter(filepath.Join(t.TempDir(), "a.sseg"), Info{}, 1)
	defer w.Abort()
	if err := w.write([]byte("abc")); err != nil || w.written != 3 {
		t.Fatal(err)
	}
	if err := brokenWriter(t, 1).write(make([]byte, 32)); err == nil {
		t.Fatal("closed file")
	}
}
