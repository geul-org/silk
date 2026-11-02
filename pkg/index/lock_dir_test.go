//ff:func feature=index type=scanner control=sequence
//ff:what TestLockDir — one writer; a dead process's lock is taken over; a path under a file fails
package index

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestLockDir(t *testing.T) {
	dir := t.TempDir()
	un, err := lockDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := lockDir(dir); !errors.Is(err, ErrLocked) {
		t.Fatal(err)
	}
	un()
	os.WriteFile(filepath.Join(dir, lockFile), []byte("0\n"), 0o644)
	un, err = lockDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	un()
	file := filepath.Join(dir, "f")
	os.WriteFile(file, nil, 0o644)
	if _, err := lockDir(filepath.Join(file, "x")); err == nil {
		t.Fatal("under a file")
	}
	ro := filepath.Join(dir, "ro")
	os.MkdirAll(filepath.Join(ro, lockFile), 0o755)
	if _, err := lockDir(filepath.Join(ro, lockFile, "..")); err == nil {
		t.Skip("running with a filesystem that allows this")
	}
}
