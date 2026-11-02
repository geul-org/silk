//ff:func feature=seg type=scanner control=sequence
//ff:what TestWriterAdd — wrong widths and out-of-order rows fail; a failed flush surfaces
package seg

import (
	"bufio"
	"errors"
	"path/filepath"
	"testing"
)

func TestWriterAdd(t *testing.T) {
	w, _ := NewWriter(filepath.Join(t.TempDir(), "a"), testInfo())
	defer w.Abort()
	if err := w.Add([]byte("x"), row(0, 0, 0)); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	if err := w.Add([]byte("pb"), row(0, 2, 0)); err != nil {
		t.Fatal(err)
	}
	if err := w.Add([]byte("pb"), row(0, 1, 0)); !errors.Is(err, ErrOrder) {
		t.Fatal(err)
	}
	if err := w.Add([]byte("pa"), row(0, 3, 0)); !errors.Is(err, ErrOrder) {
		t.Fatal(err)
	}
	w.tf.Close()
	big := make([]byte, 1<<17)
	w.tw = bufio.NewWriterSize(w.tf, 8)
	w.tw.Write(big)
	if err := w.Add([]byte("pc"), row(0, 1, 0)); err == nil {
		t.Fatal("table write on closed file")
	}
}
