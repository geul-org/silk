//ff:func feature=index type=helper control=sequence
//ff:what built — test helper: the packed test archive opened and indexed with the default spec (minimum partition 2); archive, data folder, index directory
package index

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/archive"
)

// built fails t when the build fails.
func built(t testing.TB) (*archive.Archive, string, string) {
	t.Helper()
	data := fixture.Packed(t, nil)
	a, err := archive.Open(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	dir := filepath.Join(t.TempDir(), "idx")
	if _, err := Build(context.Background(), a, dir, smallSpec(), Options{}); err != nil {
		t.Fatal(err)
	}
	return a, data, dir
}
