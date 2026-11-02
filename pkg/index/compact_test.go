//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what TestCompact — tiers merge, stale sections go major, full leaves one run; limit, lock and cancellation are honoured
package index

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/archive"
)

func TestCompact(t *testing.T) {
	ctx := context.Background()
	data, dir, a := built(t, 2)
	if _, err := Compact(ctx, dir, Options{}, false); !errors.Is(err, ErrLimit) {
		t.Fatal(err)
	}
	o := Options{Limit: 1 << 21, Policy: Policy{AppendRows: 1, RunsPerTier: 2, StaleRatio: 0.3}}
	for n := 3; n <= 6; n++ {
		fixture.Repack(t, data, docsOf(n), nil)
		a, _ = archive.Open(ctx, data)
		if _, err := Update(ctx, a, dir, o); err != nil {
			t.Fatal(err)
		}
	}
	m := current(t, dir)
	if len(m.Sections[0].Runs) < 2 {
		t.Fatal(m.Sections[0])
	}
	r, err := Compact(ctx, dir, o, false)
	if err != nil || r.Merges == 0 {
		t.Fatal(r, err)
	}
	r, err = Compact(ctx, dir, o, true)
	if err != nil {
		t.Fatal(err)
	}
	if m = current(t, dir); len(m.Sections[0].Runs) != 1 || m.Sections[0].Runs[0].Rows != 6 {
		t.Fatal(m.Sections[0])
	}
	if r, err := Compact(ctx, dir, o, true); err != nil || r.Merges != 0 {
		t.Fatal(r, err)
	}
	if err := Verify(ctx, a, dir, o); err != nil {
		t.Fatal(err)
	}
	un, _ := lockDir(dir)
	if _, err := Compact(ctx, dir, o, false); !errors.Is(err, ErrLocked) {
		t.Fatal(err)
	}
	un()
	bad := t.TempDir()
	os.WriteFile(filepath.Join(bad, currentFile), []byte("9"), 0o644)
	if _, err := Compact(ctx, bad, o, false); err == nil {
		t.Fatal("broken index")
	}
	cctx, cancel := context.WithCancel(ctx)
	cancel()
	if _, err := Compact(cctx, dir, o, false); !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}
