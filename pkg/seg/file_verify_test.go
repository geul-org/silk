//ff:func feature=seg type=scanner control=sequence
//ff:what TestFileVerify — a sound segment passes; a flipped byte, broken rows and a broken table fail
package seg

import (
	"crypto/sha256"
	"errors"
	"slices"
	"testing"
)

func TestFileVerify(t *testing.T) {
	f := writeSeg(t, "a")
	if err := f.Verify(); err != nil {
		t.Fatal(err)
	}
	fix := func(g *File) {
		e := uint64(len(g.b) - footerSize)
		s := sha256.Sum256(g.b[:e])
		copy(g.b[e+64:], s[:])
	}
	damage := func(mut func(b []byte)) error {
		b := slices.Clone(f.b)
		mut(b)
		g, err := frame("x", b)
		if err != nil {
			t.Fatal(err)
		}
		fix(g)
		return g.Verify()
	}
	if err := damage(func(b []byte) { copy(b, row(9, 9, 9)) }); !errors.Is(err, ErrFormat) {
		t.Fatal("rows", err)
	}
	if err := damage(func(b []byte) { copy(b[18:], "pb") }); !errors.Is(err, ErrFormat) {
		t.Fatal("table", err)
	}
	if err := damage(func(b []byte) { b[18+2+16+2+15] = 0 }); !errors.Is(err, ErrFormat) {
		t.Fatal("count", err)
	}
	b := slices.Clone(f.b)
	b[0] ^= 1
	g, _ := frame("x", b)
	if err := g.Verify(); !errors.Is(err, ErrFormat) {
		t.Fatal("hash", err)
	}
}
