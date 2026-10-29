//ff:func feature=index type=builder control=sequence
//ff:what TestBuild — a build in small batches merges to one run per section and verifies; rebuilding keeps file numbers rising; bad spec, lock, archive, documents, segment and table writes fail
package index

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/spec"
)

func TestBuild(t *testing.T) {
	ctx := context.Background()
	data := fixture.Packed(t, nil)
	a, _ := archive.Open(ctx, data)
	dir := filepath.Join(t.TempDir(), "idx")
	r, err := Build(ctx, a, dir, smallSpec(), Options{Batch: 50})
	m, _ := CurrentManifest(dir)
	if err != nil || r.Docs != 4 || r.New != 4 || r.NewBytes == 0 || r.Written == 0 || len(m.Sections[0].Runs) != 1 || len(r.Tasks) == 0 {
		t.Fatal(err, r)
	}
	if _, err := Verify(ctx, a, dir); err != nil {
		t.Fatal(err)
	}
	next := m.NextFile
	Build(ctx, a, dir, smallSpec(), Options{})
	if m2, _ := CurrentManifest(dir); m2.NextFile <= next || m2.Generation <= m.Generation {
		t.Fatal(m2.NextFile, next)
	}
	if _, err := Build(ctx, a, dir, spec.Spec{}, Options{}); !errors.Is(err, spec.ErrSpec) {
		t.Fatal(err)
	}
	release, _ := lockDir(dir)
	if _, err := Build(ctx, a, dir, smallSpec(), Options{}); !errors.Is(err, ErrLocked) {
		t.Fatal(err)
	}
	release()
	failBuild := func(name string, prep func(d string), o Options, ar *archive.Archive) {
		d := filepath.Join(t.TempDir(), "idx")
		prep(d)
		if _, err := Build(ctx, ar, d, smallSpec(), o); err == nil {
			t.Fatal(name)
		}
	}
	failBuild("segments while adding", func(d string) { block(d, dirSeg) }, Options{Batch: 1}, a)
	failBuild("segments at the end", func(d string) { block(d, dirSeg) }, Options{}, a)
	failBuild("document table", func(d string) { block(d, dirDoc) }, Options{}, a)
	b, _ := archive.Open(ctx, data)
	os.Remove(filepath.Join(b.Sets[0].Dir, "m", b.Sets[0].Manifest+".gpm"))
	failBuild("archive", func(string) {}, Options{}, b)
}
