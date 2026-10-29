//ff:func feature=index type=builder control=sequence
//ff:what TestUpdate — the evolved archive adds appended segments and dead rows and verifies; nothing new is Unchanged; a spec change rebuilds its tag; bad spec, lock, index, tables and archive fail
package index

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/spec"
)

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	_, data, dir := built(t)
	a := evolved(t, data)
	r, err := Update(ctx, a, dir, smallSpec(), Options{})
	m, _ := CurrentManifest(dir)
	if err != nil || r.New != 2 || r.Killed != 2 || r.Kept != 2 || r.NewBytes == 0 || r.Read != r.NewBytes || len(m.Sections[0].Appends) != 1 || m.Sections[0].Dead == 0 {
		t.Fatal(err, r)
	}
	if _, err := Verify(ctx, a, dir); err != nil {
		t.Fatal(err)
	}
	if r, err := Update(ctx, a, dir, smallSpec(), Options{}); err != nil || !r.Unchanged || r.Docs != 4 {
		t.Fatal(err, r)
	}
	sp := smallSpec()
	sp.Tags[2].Sort.Width = 8
	if r, err := Update(ctx, a, dir, sp, Options{Compact: true}); err != nil || r.Read == 0 || r.New != 0 {
		t.Fatal(err, r)
	}
	if _, err := Verify(ctx, a, dir); err != nil {
		t.Fatal(err)
	}
	if _, err := Update(ctx, a, dir, spec.Spec{}, Options{}); !errors.Is(err, spec.ErrSpec) {
		t.Fatal(err)
	}
	release, _ := lockDir(dir)
	if _, err := Update(ctx, a, dir, sp, Options{}); !errors.Is(err, ErrLocked) {
		t.Fatal(err)
	}
	release()
	if _, err := Update(ctx, a, t.TempDir(), sp, Options{}); !errors.Is(err, ErrNoIndex) {
		t.Fatal(err)
	}
	failUpdate := func(name string, prep func(dir string, m Manifest) *archive.Archive, o Options) {
		_, data, dir := built(t)
		b := evolved(t, data)
		m, _ := CurrentManifest(dir)
		if p := prep(dir, m); p != nil {
			b = p
		}
		if _, err := Update(ctx, b, dir, smallSpec(), o); err == nil {
			t.Fatal(name)
		}
	}
	failUpdate("document table", func(d string, m Manifest) *archive.Archive { os.Remove(filepath.Join(d, m.Docs[0].File)); return nil }, Options{})
	failUpdate("bitmap", func(d string, m Manifest) *archive.Archive { os.Remove(filepath.Join(d, m.Live)); return nil }, Options{})
	failUpdate("segments", func(d string, m Manifest) *archive.Archive { block(d, dirSeg); return nil }, Options{})
	failUpdate("archive", func(d string, m Manifest) *archive.Archive {
		b, _ := archive.Open(ctx, data)
		os.Remove(filepath.Join(b.Sets[0].Dir, "m", b.Sets[0].Manifest+".gpm"))
		return b
	}, Options{})
}
