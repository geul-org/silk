//ff:func feature=seg type=reader control=sequence
//ff:what TestOpen — a written segment opens with path and size; a missing, unreadable or malformed file fails
package seg

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpen(t *testing.T) {
	f := sampleSeg(t, true)
	if f.Info.Bytes == 0 || len(f.Parts) != 2 || f.Info.Other != 2 {
		t.Fatal(f.Info, f.Parts)
	}
	d := t.TempDir()
	if _, err := Open(filepath.Join(d, "none")); err == nil {
		t.Fatal("missing")
	}
	bad := filepath.Join(d, "bad")
	os.WriteFile(bad, []byte("short"), 0o644)
	if _, err := Open(bad); err == nil {
		t.Fatal("malformed")
	}
	os.Chmod(bad, 0)
	if _, err := os.ReadFile(bad); err == nil {
		t.Skip("running as root: permissions not enforced")
	}
	if _, err := Open(bad); err == nil {
		t.Fatal("unreadable")
	}
}
