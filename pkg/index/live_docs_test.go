//ff:func feature=index type=reader control=sequence
//ff:what TestLiveDocs — the four test documents with set, pack and place; an unreadable set manifest fails
package index

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestLiveDocs(t *testing.T) {
	a, _, _ := built(t)
	ds, err := liveDocs(context.Background(), a)
	if err != nil || len(ds) != 4 || ds[0].Set != a.Sets[0].Name || ds[0].Doc.Len == 0 || ds[0].Path == "" {
		t.Fatal(err, ds)
	}
	os.Remove(filepath.Join(a.Sets[0].Dir, "m", a.Sets[0].Manifest+".gpm"))
	if _, err := liveDocs(context.Background(), a); err == nil {
		t.Fatal("missing manifest")
	}
}
