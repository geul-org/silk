//ff:func feature=seg type=helper control=sequence
//ff:what TestWriterWritten — counts bytes handed to the file
package seg

import (
	"path/filepath"
	"testing"
)

func TestWriterWritten(t *testing.T) {
	w, _ := NewWriter(filepath.Join(t.TempDir(), "a.sseg"), Info{}, 1)
	defer w.Abort()
	w.write(make([]byte, 5))
	if w.Written() != 5 {
		t.Fatal(w.Written())
	}
}
