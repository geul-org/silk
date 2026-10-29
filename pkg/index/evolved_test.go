//ff:func feature=index type=helper control=sequence
//ff:what evolved — test helper: the data folder's next generation (one document changed, one removed, one added), opened
package index

import (
	"context"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/archive"
)

// evolved runs fixture.Evolve and reopens the archive.
func evolved(t testing.TB, data string) *archive.Archive {
	t.Helper()
	fixture.Evolve(t, data)
	a, err := archive.Open(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	return a
}
