//ff:func feature=index type=helper control=sequence
//ff:what setup — a data folder of n documents, an index directory with an INDEX and a SILK declared, and the archive
package index

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/decl"
)

func setup(t *testing.T, n int) (string, string, *archive.Archive) {
	t.Helper()
	data := fixture.Pack(t, docsOf(n), nil)
	dir := filepath.Join(t.TempDir(), "ix")
	g := fixture.Grammar(t)
	d, err := decl.NewIndex(g, "t", "TRIPLE", "predicate", []string{"object"}, []int{6})
	if err != nil {
		t.Fatal(err)
	}
	if err := Declare(dir, d, g); err != nil {
		t.Fatal(err)
	}
	s, _ := decl.NewSilk("s", silkDef)
	if err := Declare(dir, s, g); err != nil {
		t.Fatal(err)
	}
	a, err := archive.Open(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	return data, dir, a
}
