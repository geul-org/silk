//ff:func feature=index type=writer control=sequence
//ff:what TestCompact — due merges run and commit a new generation that still verifies; a held lock or no index fails
package index

import (
	"context"
	"errors"
	"testing"
)

func TestCompact(t *testing.T) {
	a, _, dir := updated(t)
	ctx := context.Background()
	r, err := Compact(ctx, dir, Options{Full: true})
	m, _ := CurrentManifest(dir)
	if err != nil || len(r.Tasks) != 7 || r.Generation != m.Generation || m.Sections[0].Dead != 0 || len(m.Sections[0].Runs) != 1 || r.Elapsed == 0 {
		t.Fatal(err, r)
	}
	if _, err := Verify(ctx, a, dir); err != nil {
		t.Fatal(err)
	}
	release, _ := lockDir(dir)
	if _, err := Compact(ctx, dir, Options{}); !errors.Is(err, ErrLocked) {
		t.Fatal(err)
	}
	release()
	if _, err := Compact(ctx, t.TempDir(), Options{}); !errors.Is(err, ErrNoIndex) {
		t.Fatal(err)
	}
}
