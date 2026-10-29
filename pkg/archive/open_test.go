//ff:func feature=archive type=reader control=sequence
//ff:what TestOpen — a data folder and its .well-known/geul open to the same set; a plain tree is ErrNotPacked; a bad latest.txt fails
package archive

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestOpen(t *testing.T) {
	ctx := context.Background()
	dir := fixture.Packed(t, nil)
	a, err := Open(ctx, dir)
	if err != nil || len(a.Sets) != 1 || a.Sets[0].Name != "0x00010000" || a.Sets[0].Generation != 1 || len(a.Sets[0].Manifest) != 64 {
		t.Fatal(a, err)
	}
	wk := filepath.Join(dir, ".well-known", "geul")
	if b, err := Open(ctx, wk); err != nil || b.Root != wk || b.Sets[0] != a.Sets[0] {
		t.Fatal(b, err)
	}
	if _, err := Open(ctx, fixture.Source()); !errors.Is(err, ErrNotPacked) {
		t.Fatal(err)
	}
	os.WriteFile(filepath.Join(wk, "0x00010000", "latest.txt"), []byte("garbage"), 0o644)
	if _, err := Open(ctx, dir); err == nil {
		t.Fatal("bad latest.txt must fail")
	}
}
