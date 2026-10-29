//ff:func feature=index type=builder control=sequence
//ff:what TestApplyDiff — new versions get sequence numbers after the old ones, killed ones die with their rows counted dead; failing kept, added or flushed segments fail
package index

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/silk/pkg/archive"
)

func TestApplyDiff(t *testing.T) {
	ctx := context.Background()
	var arch *archive.Archive
	prep := func() (string, Manifest, []DocEntry, Bitmap, docDiff) {
		_, data, dir := built(t)
		arch = evolved(t, data)
		m, _ := CurrentManifest(dir)
		old, _ := loadDocs(dir, m)
		live, _ := readBitmap(filepath.Join(dir, m.Live))
		docs, _ := liveDocs(ctx, arch)
		return dir, m, old, live, diffDocs(m, old, live, docs)
	}
	dir, m, old, live, d := prep()
	r, err := applyDiff(dir, &m, smallSpec(), arch, old, live, d, fill(7, false), Options{})
	if err != nil || r.New != 2 || r.Killed != 2 || r.Docs != 4 || m.Docs[1].First != 4 || m.Sections[0].Dead == 0 {
		t.Fatal(err, r)
	}
	rebuild := fill(7, false)
	rebuild[0] = true
	dir, m, old, live, d = prep()
	block(dir, dirSeg)
	if _, err := applyDiff(dir, &m, smallSpec(), arch, old, live, d, rebuild, Options{Batch: 1}); err == nil {
		t.Fatal("kept")
	}
	dir, m, old, live, d = prep()
	block(dir, dirSeg)
	if _, err := applyDiff(dir, &m, smallSpec(), arch, old, live, d, fill(7, false), Options{Batch: 1}); err == nil {
		t.Fatal("added")
	}
	dir, m, old, live, d = prep()
	block(dir, dirSeg)
	if _, err := applyDiff(dir, &m, smallSpec(), arch, old, live, d, fill(7, false), Options{}); err == nil {
		t.Fatal("flush")
	}
}
