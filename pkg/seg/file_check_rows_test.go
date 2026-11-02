//ff:func feature=seg type=scanner control=sequence
//ff:what TestFileCheckRows — rows out of order in a partition are reported
package seg

import (
	"errors"
	"slices"
	"testing"
)

func TestFileCheckRows(t *testing.T) {
	f := writeSeg(t, "a")
	if err := f.checkRows(0); err != nil {
		t.Fatal(err)
	}
	b := slices.Clone(f.b)
	copy(b, row(0, 9, 1))
	g, _ := frame("x", b)
	if err := g.checkRows(0); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
}
