//ff:func feature=index type=writer control=sequence
//ff:what TestLockDir — one writer at a time; a dead holder's lock is taken over; a blocked or unwritable directory fails
package index

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestLockDir(t *testing.T) {
	d := filepath.Join(t.TempDir(), "idx")
	release, err := lockDir(d)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := lockDir(d); !errors.Is(err, ErrLocked) {
		t.Fatal(err)
	}
	release()
	os.WriteFile(filepath.Join(d, lockFile), fmt.Appendf(nil, "%d\n", 1<<30), 0o644)
	if r, err := lockDir(d); err != nil {
		t.Fatal("stale lock", err)
	} else {
		r()
	}
	e := t.TempDir()
	block(e, "x")
	if _, err := lockDir(filepath.Join(e, "x")); err == nil {
		t.Fatal("blocked directory")
	}
	ro := filepath.Join(t.TempDir(), "ro")
	os.Mkdir(ro, 0o500)
	if _, err := lockDir(ro); err == nil && os.Getuid() != 0 {
		t.Fatal("unwritable directory")
	}
}
