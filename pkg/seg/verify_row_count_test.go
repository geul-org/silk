//ff:func feature=seg type=scanner control=sequence
//ff:what TestVerifyRowCount — partitions that do not reach the footer's row count fail
package seg

import (
	"crypto/sha256"
	"errors"
	"slices"
	"testing"
)

func TestVerifyRowCount(t *testing.T) {
	f := writeSeg(t, "a")
	b := slices.Clone(f.b)
	g, err := frame("x", b)
	if err != nil {
		t.Fatal(err)
	}
	g.Parts = 1
	g.table = g.table[:18]
	e := uint64(len(g.b) - footerSize)
	s := sha256.Sum256(g.b[:e])
	copy(g.b[e+64:], s[:])
	if err := g.Verify(); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
}
