//ff:func feature=index type=helper control=sequence
//ff:what built — setup, then a build
package index

import (
	"context"
	"testing"

	"github.com/geul-org/silk/pkg/archive"
)

func built(t *testing.T, n int) (string, string, *archive.Archive) {
	t.Helper()
	data, dir, a := setup(t, n)
	if _, err := Build(context.Background(), a, dir, testOpts); err != nil {
		t.Fatal(err)
	}
	return data, dir, a
}
