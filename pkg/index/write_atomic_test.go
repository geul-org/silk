//ff:func feature=index type=writer control=sequence
//ff:what TestWriteAtomic — bytes land at path; a blocked directory, an unwritable temporary path and a full device fail
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteAtomic(t *testing.T) {
	d := t.TempDir()
	p := filepath.Join(d, "a", "f")
	if n, err := writeAtomic(p, []byte("xy")); err != nil || n != 2 {
		t.Fatal(err)
	}
	if b, _ := os.ReadFile(p); string(b) != "xy" {
		t.Fatal(string(b))
	}
	block(d, "b")
	if _, err := writeAtomic(filepath.Join(d, "b", "f"), nil); err == nil {
		t.Fatal("blocked directory")
	}
	os.Mkdir(filepath.Join(d, "c.tmp"), 0o755)
	if _, err := writeAtomic(filepath.Join(d, "c"), nil); err == nil {
		t.Fatal("temporary path is a directory")
	}
	if _, err := os.Stat("/dev/full"); err != nil {
		t.Skip("no /dev/full")
	}
	os.Symlink("/dev/full", filepath.Join(d, "full.tmp"))
	if _, err := writeAtomic(filepath.Join(d, "full"), []byte("x")); err == nil {
		t.Fatal("full device")
	}
}
