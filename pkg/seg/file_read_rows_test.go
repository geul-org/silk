//ff:func feature=seg type=reader control=sequence
//ff:what TestFileReadRows — rows of a region by index; a failing read is ErrFormat
package seg

import (
	"bytes"
	"errors"
	"testing"
)

func TestFileReadRows(t *testing.T) {
	f := sampleSeg(t, true)
	b, err := f.readRows(f.partRegion(0), 1, 2)
	if err != nil || !bytes.Equal(b, append(row("x1", 2, 4), row("x2", 3, 9)...)) {
		t.Fatalf("%v %x", err, b)
	}
	if _, err := failing(f, 0).readRows(f.partRegion(0), 0, 1); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
}
