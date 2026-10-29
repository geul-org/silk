//ff:func feature=seg type=writer control=sequence
//ff:what TestWriterEnd — an empty partition leaves nothing, a full one a table entry, a small one rows in the other partition behind its key
package seg

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestWriterEnd(t *testing.T) {
	w, _ := NewWriter(filepath.Join(t.TempDir(), "a.sseg"), Info{PartWidth: 2, SortWidth: 2}, 2)
	defer w.Abort()
	w.Begin([]byte("aa"))
	w.End()
	w.Begin([]byte("bb"))
	w.Add(row("k1", 1, 0))
	w.Add(row("k2", 1, 2))
	w.End()
	w.Begin([]byte("cc"))
	w.Add(row("k3", 2, 0))
	w.End()
	if len(w.parts) != 1 || w.parts[0].Count != 2 || w.info.Rows != 2 || w.info.Other != 1 || !bytes.Equal(w.other, append([]byte("cc"), row("k3", 2, 0)...)) {
		t.Fatal(w.parts, w.info, w.other)
	}
}
