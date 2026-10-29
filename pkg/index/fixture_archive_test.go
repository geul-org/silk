//ff:func feature=index type=helper control=sequence
//ff:what fixtureArchive — the raw test archive opened, and a fresh index directory path
package index

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/silk/internal/fixture"
	"github.com/park-jun-woo/silk/pkg/archive"
)

// fixtureArchive packs testdata/archive and opens it.
func fixtureArchive(t *testing.T) (*archive.Archive, string) {
	t.Helper()
	a, err := archive.Open(context.Background(), fixture.Packed(t, nil))
	if err != nil {
		t.Fatal(err)
	}
	return a, filepath.Join(t.TempDir(), "idx")
}
