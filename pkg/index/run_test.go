//ff:func feature=index type=scanner control=sequence
//ff:what TestRun — a limit is required, the lock holds, a damaged index fails, build resets and update keeps
package index

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestRun(t *testing.T) {
	ctx := context.Background()
	_, dir, a := setup(t, 3)
	if _, err := run(ctx, a, dir, Options{}, true); !errors.Is(err, ErrLimit) {
		t.Fatal(err)
	}
	un, _ := lockDir(dir)
	if _, err := run(ctx, a, dir, testOpts, false); !errors.Is(err, ErrLocked) {
		t.Fatal(err)
	}
	un()
	r, err := run(ctx, a, dir, testOpts, false)
	if err != nil || r.NewDocs != 3 || r.ReadDocs != 3 {
		t.Fatal(r, err)
	}
	r, err = Update(ctx, a, dir, testOpts)
	if err != nil || r.NewDocs != 0 || r.ReadDocs != 0 {
		t.Fatal(r, err)
	}
	r, err = run(ctx, a, dir, testOpts, true)
	if err != nil || r.NewDocs != 3 {
		t.Fatal(r, err)
	}
	m := current(t, dir)
	if out := resetManifest(m); len(out.Docs) != 0 || !out.Sections[0].Pending || out.NextFile != m.NextFile {
		t.Fatal(out)
	}
	if rs := setRefs(a); len(rs) != 1 || !slices.Equal(rs, m.Sets) {
		t.Fatal(rs)
	}
	os.Remove(filepath.Join(dir, m.Live))
	if _, err := run(ctx, a, dir, testOpts, false); err == nil {
		t.Fatal("missing live")
	}
	os.WriteFile(filepath.Join(dir, currentFile), []byte("x"), 0o644)
	if _, err := run(ctx, a, dir, testOpts, false); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
}
