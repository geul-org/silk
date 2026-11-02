//ff:func feature=index type=scanner control=sequence
//ff:what TestPins — pinning, retrying when CURRENT moves, parsing and pruning pins
package index

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestPins(t *testing.T) {
	dir := t.TempDir()
	if _, _, err := pinCurrent(dir, readCurrent); !errors.Is(err, ErrNoIndex) {
		t.Fatal(err)
	}
	n := uint64(0)
	moving := func(string) (uint64, error) { n++; return n, nil }
	if _, _, err := pinCurrent(dir, moving); !errors.Is(err, ErrLocked) {
		t.Fatal(err)
	}
	same := func(string) (uint64, error) { return 5, nil }
	g, p, err := pinCurrent(dir, same)
	if err != nil || g != 5 {
		t.Fatal(g, err)
	}
	if !pinnedGens(dir)[5] {
		t.Fatal("pinned")
	}
	os.Remove(p)
	os.WriteFile(filepath.Join(dir, dirPin, "bad"), nil, 0o644)
	os.WriteFile(filepath.Join(dir, dirPin, "7-0-1"), nil, 0o644)
	if gs := pinnedGens(dir); len(gs) != 0 {
		t.Fatal(gs)
	}
	if es, _ := os.ReadDir(filepath.Join(dir, dirPin)); len(es) != 0 {
		t.Fatal("left over", es)
	}
	if _, _, ok := parsePin("1-2"); ok {
		t.Fatal("short")
	}
	file := filepath.Join(dir, "f")
	os.WriteFile(file, nil, 0o644)
	if _, err := pinGen(file, 1); err == nil {
		t.Fatal("under a file")
	}
	if _, _, err := pinCurrent(file, same); err == nil {
		t.Fatal("pin failure")
	}
}
