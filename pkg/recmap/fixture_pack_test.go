//ff:func feature=recmap type=helper control=sequence
//ff:what fixturePack — the raw test archive's one data pack, its file open for reading (closed at test end)
package recmap

import (
	"context"
	"os"
	"testing"

	"github.com/park-jun-woo/silk/internal/fixture"
	"github.com/park-jun-woo/silk/pkg/archive"
)

// fixturePack packs testdata/archive and returns its data pack.
func fixturePack(t *testing.T) (archive.Pack, *os.File) {
	t.Helper()
	a, err := archive.Open(context.Background(), fixture.Packed(t, nil))
	if err != nil {
		t.Fatal(err)
	}
	ps, err := a.Sets[0].Packs(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	f, err := os.Open(ps[0].Path)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	return ps[0], f
}
