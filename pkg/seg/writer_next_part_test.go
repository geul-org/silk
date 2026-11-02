//ff:func feature=seg type=scanner control=sequence
//ff:what TestWriterNextPart — closing a partition writes its table entry; failure stops
package seg

import (
	"bufio"
	"path/filepath"
	"testing"
)

func TestWriterNextPart(t *testing.T) {
	w, _ := NewWriter(filepath.Join(t.TempDir(), "a"), testInfo())
	defer w.Abort()
	if err := w.closePart(); err != nil || w.parts != 0 {
		t.Fatal(err)
	}
	if err := w.nextPart([]byte("pa")); err != nil || !w.open {
		t.Fatal(err)
	}
	w.rows = 2
	if err := w.nextPart([]byte("pb")); err != nil || w.parts != 1 || w.start != 2 {
		t.Fatal(err, w.parts)
	}
	w.tw = bufio.NewWriterSize(failWriter{}, 8)
	if err := w.nextPart([]byte("pc")); err == nil {
		t.Fatal("no error")
	}
}
