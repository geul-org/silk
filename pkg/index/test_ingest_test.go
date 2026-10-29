//ff:func feature=index type=helper control=sequence
//ff:what testIngest — test helper: an ingest over a fresh index directory with the small spec, every tag sorted, and the archive's live documents
package index

import (
	"context"
	"path/filepath"
	"testing"
)

// testIngest returns the ingest, its manifest, directory and documents.
func testIngest(t *testing.T, batch uint64) (*ingest, *Manifest, string, []liveDoc) {
	t.Helper()
	a, _, _ := built(t)
	docs, err := liveDocs(context.Background(), a)
	if err != nil {
		t.Fatal(err)
	}
	m := &Manifest{}
	dir := filepath.Join(t.TempDir(), "g")
	sp := smallSpec()
	return newIngest(dir, m, sp, fill(len(sp.Tags), true), batch), m, dir, docs
}
