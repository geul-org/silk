//ff:func feature=index type=scanner control=sequence
//ff:what TestLoadSectionsPending — a declared but unbuilt section is known and has no files
package index

import (
	"context"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/decl"
)

func TestLoadSectionsPending(t *testing.T) {
	ctx := context.Background()
	data, dir, a := built(t, 1)
	s, _ := decl.NewSilk("later", silkDef+"# later\n")
	if err := Declare(dir, s, fixture.Grammar(t)); err != nil {
		t.Fatal(err)
	}
	a, _ = archive.Open(ctx, data)
	ix, err := Open(ctx, a, dir, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer ix.Close()
	if ix.Section("later", 0) != nil || ix.Decls["later"].Name != "later" {
		t.Fatal("pending")
	}
	if _, err := Update(ctx, a, dir, testOpts); err != nil {
		t.Fatal(err)
	}
	ix2, err := Open(ctx, a, dir, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer ix2.Close()
	if ix2.Section("later", 0) == nil {
		t.Fatal("built")
	}
}
