//ff:func feature=query type=helper control=iteration dimension=1
//ff:what testIndex — test helper: the packed test archive indexed with the default spec at a given minimum partition, opened (closed at cleanup)
package query

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/silk/internal/fixture"
	"github.com/park-jun-woo/silk/pkg/archive"
	"github.com/park-jun-woo/silk/pkg/index"
	"github.com/park-jun-woo/silk/pkg/spec"
)

// testIndex uses min for every tag (1: table partitions, 256: other partitions).
func testIndex(t testing.TB, min int) *index.Index {
	t.Helper()
	ctx := context.Background()
	a, err := archive.Open(ctx, fixture.Packed(t, nil))
	if err != nil {
		t.Fatal(err)
	}
	sp := spec.Default()
	for i := range sp.Tags {
		sp.Tags[i].MinPartition = min
	}
	dir := filepath.Join(t.TempDir(), "idx")
	if _, err := index.Build(ctx, a, dir, sp, index.Options{}); err != nil {
		t.Fatal(err)
	}
	ix, err := index.Open(a, dir)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { ix.Close() })
	return ix
}
