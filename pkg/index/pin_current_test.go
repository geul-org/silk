//ff:func feature=index type=reader control=sequence
//ff:what TestPinCurrent — pins the stable generation; a failing read, a blocked pin directory or a CURRENT that keeps changing fails
package index

import (
	"errors"
	"testing"
)

func TestPinCurrent(t *testing.T) {
	d := t.TempDir()
	fixed := func(string) (uint64, error) { return 3, nil }
	if g, p, err := pinCurrent(d, fixed); err != nil || g != 3 || p == "" {
		t.Fatal(g, p, err)
	}
	if _, _, err := pinCurrent(d, readCurrent); !errors.Is(err, ErrNoIndex) {
		t.Fatal(err)
	}
	n := uint64(0)
	moving := func(string) (uint64, error) { n++; return n, nil }
	if _, _, err := pinCurrent(d, moving); !errors.Is(err, ErrLocked) {
		t.Fatal(err)
	}
	if pinnedGens(d)[1] {
		t.Fatal("pins of a moving CURRENT must be removed")
	}
	e := t.TempDir()
	block(e, dirPin)
	if _, _, err := pinCurrent(e, fixed); err == nil {
		t.Fatal("blocked pin directory")
	}
}
