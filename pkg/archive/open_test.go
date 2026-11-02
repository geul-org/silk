//ff:func feature=archive type=scanner control=sequence
//ff:what TestOpen — a data folder and its well-known directory open alike; other directories and broken sets fail
package archive

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestOpen(t *testing.T) {
	data, a := testArchive(t)
	if len(a.Sets) != 1 || a.Set("0x00010000") == nil || a.Set("nope") != nil {
		t.Fatal(a.Sets)
	}
	if b, err := Open(context.Background(), wellKnown(data)); err != nil || b.Root != a.Root {
		t.Fatal(err)
	}
	if _, err := Open(context.Background(), t.TempDir()); !errors.Is(err, ErrNotPacked) {
		t.Fatal(err)
	}
	s := a.Sets[0]
	os.WriteFile(filepath.Join(s.Dir, "latest.txt"), []byte("bad"), 0o644)
	if _, err := Open(context.Background(), data); err == nil {
		t.Fatal("bad latest")
	}
}
