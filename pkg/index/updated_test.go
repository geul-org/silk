//ff:func feature=index type=helper control=sequence
//ff:what updated — test helper: a built index brought to the evolved archive by Update (appended segments and dead rows in every section); archive, data folder, index directory
package index

import (
	"context"
	"testing"

	"github.com/geul-org/silk/pkg/archive"
)

// updated fails t when the update fails.
func updated(t testing.TB) (*archive.Archive, string, string) {
	t.Helper()
	_, data, dir := built(t)
	a := evolved(t, data)
	if _, err := Update(context.Background(), a, dir, smallSpec(), Options{}); err != nil {
		t.Fatal(err)
	}
	return a, data, dir
}
