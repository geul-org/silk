//ff:func feature=query type=helper control=sequence
//ff:what fixtureIndex — the raw test archive indexed and opened (closed at test end)
package query

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/silk/internal/fixture"
	"github.com/park-jun-woo/silk/pkg/archive"
	"github.com/park-jun-woo/silk/pkg/index"
)

// fixtureIndex builds and opens the index of testdata/archive.
func fixtureIndex(t *testing.T) *index.Index {
	t.Helper()
	ctx := context.Background()
	a, err := archive.Open(ctx, fixture.Packed(t, nil))
	if err != nil {
		t.Fatal(err)
	}
	dir := filepath.Join(t.TempDir(), "idx")
	if _, err := index.Build(ctx, a, dir); err != nil {
		t.Fatal(err)
	}
	ix, err := index.Open(a, dir)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { ix.Close() })
	return ix
}
