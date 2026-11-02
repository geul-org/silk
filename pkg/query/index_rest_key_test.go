//ff:func feature=query type=scanner control=iteration dimension=1
//ff:what TestIndexRestKey — an INDEX key on text bytes finds them, numbered and named paths alike, and is covered only when the text fits its slot
package query

import (
	"context"
	"path/filepath"
	"slices"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/index"
)

func TestIndexRestKey(t *testing.T) {
	ctx := context.Background()
	g := fixture.Grammar(t)
	dir := filepath.Join(t.TempDir(), "ix")
	d, err := decl.NewIndex(g, "l", "LABEL", "object.language", []string{"object.bytes"}, []int{8})
	if err != nil {
		t.Fatal(err)
	}
	if err := index.Declare(dir, d, g); err != nil {
		t.Fatal(err)
	}
	data := fixture.Pack(t, fixture.Corpus(9, 400), nil)
	a, _ := archive.Open(ctx, data)
	if _, err := index.Build(ctx, a, dir, testOpts); err != nil {
		t.Fatal(err)
	}
	ix, err := index.Open(ctx, a, dir, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer ix.Close()
	for _, w := range [][]string{{"LABEL", "object.language=@en", `object.bytes="Par"`}, {"LABEL", "2.0=@en", `2.1="Par"`}, {"LABEL", "object.language=@en", `object.bytes="Abcdefghij"`}} {
		if x, y := answer(t, ix, false, w...), answer(t, ix, true, w...); !slices.Equal(x, y) || len(x) == 0 {
			t.Fatal(w, x, y)
		}
	}
	q, _ := Parse(ix, []string{"LABEL", "object.language=@en", `2.1="Par"`})
	p, _ := makePlan(ix, &q.Conds[0], "")
	if !p.Covered {
		t.Fatal(p.describe())
	}
	q, _ = Parse(ix, []string{"LABEL", "object.language=@en", `2.1="Abcdefghij"`})
	if p, _ = makePlan(ix, &q.Conds[0], ""); p.Covered {
		t.Fatal(p.describe())
	}
}
