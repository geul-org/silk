//ff:func feature=index type=scanner control=sequence
//ff:what TestLockCreateFails — a lock path that is a directory cannot be created
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLockCreateFails(t *testing.T) {
	dir := t.TempDir()
	os.Mkdir(filepath.Join(dir, lockFile), 0o755)
	os.WriteFile(filepath.Join(dir, lockFile, "x"), nil, 0o644)
	if _, err := lockDir(dir); err == nil {
		t.Fatal("LOCK is a full directory")
	}
}
