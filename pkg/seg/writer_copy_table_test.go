//ff:func feature=seg type=scanner control=sequence
//ff:what TestWriterCopyTable — a closed table file fails at seek and a bad destination at copy
package seg

import (
	"bufio"
	"path/filepath"
	"testing"
)

func TestWriterCopyTable(t *testing.T) {
	w, _ := NewWriter(filepath.Join(t.TempDir(), "a"), testInfo())
	defer w.Abort()
	w.tf.WriteString("data")
	w.bw = bufio.NewWriterSize(failWriter{}, 1)
	if err := w.copyTable(); err == nil {
		t.Fatal("copy to a failing writer")
	}
	w.tf.Close()
	if err := w.copyTable(); err == nil {
		t.Fatal("seek on closed file")
	}
}
