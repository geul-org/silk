//ff:func feature=query type=helper control=sequence
//ff:what openIndex — a built index over a data folder
package query

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/index"
)

func openIndex(t *testing.T, data string) (*archive.Archive, string, *index.Index) {
	t.Helper()
	ctx := context.Background()
	dir := filepath.Join(t.TempDir(), "ix")
	declareAll(t, dir)
	a, err := archive.Open(ctx, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := index.Build(ctx, a, dir, testOpts); err != nil {
		t.Fatal(err)
	}
	ix, err := index.Open(ctx, a, dir, 1<<20)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(ix.Close)
	return a, dir, ix
}
