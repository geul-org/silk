//ff:func feature=index type=rule control=sequence
//ff:what TestDiffDocs — against the evolved archive: two kept, the rewritten and the added version new, the rewritten and the removed version killed; dead entries are ignored
package index

import (
	"context"
	"path/filepath"
	"testing"
)

func TestDiffDocs(t *testing.T) {
	_, data, dir := built(t)
	m, _ := CurrentManifest(dir)
	old, _ := loadDocs(dir, m)
	live, _ := readBitmap(filepath.Join(dir, m.Live))
	docs, _ := liveDocs(context.Background(), evolved(t, data))
	d := diffDocs(m, old, live, docs)
	if len(d.kept) != 2 || len(d.added) != 2 || len(d.killed) != 2 || d.killed[0] > d.killed[1] {
		t.Fatal(d)
	}
	live.Put(d.kept[0].seq, false)
	if d2 := diffDocs(m, old, live, docs); len(d2.kept) != 1 || len(d2.added) != 3 || len(d2.killed) != 2 {
		t.Fatal(d2)
	}
}
