//ff:func feature=fixture type=builder control=sequence
//ff:what TestEvolve — a packed folder gains generation 2; an unwritable source copy and a failing pack write report through t.Fatal
package fixture

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestEvolve(t *testing.T) {
	out := Packed(t, nil)
	Evolve(t, out)
	b, err := os.ReadFile(filepath.Join(out, ".well-known", "geul", "0x00010000", "latest.txt"))
	if err != nil || !strings.HasPrefix(string(b), "2 ") {
		t.Fatal(string(b), err)
	}
	r := &recorder{T: t}
	blocked := t.TempDir()
	os.WriteFile(filepath.Join(blocked, ".well-known"), nil, 0o644)
	Evolve(r, blocked)
	if !r.failed {
		t.Fatal("a failing pack write must call Fatal")
	}
	r2 := &recorder{T: t, tmp: filepath.Join(blocked, ".well-known", "x")}
	Evolve(r2, out)
	if !r2.failed {
		t.Fatal("a failing copy must call Fatal")
	}
}
