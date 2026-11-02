//ff:func feature=query type=scanner control=sequence
//ff:what TestUpdateCompact — after an update, and after each compaction, answers equal a fresh build's; a full compaction gives its rows byte for byte
package query

import (
	"context"
	"errors"
	"math/rand"
	"slices"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/index"
)

func TestUpdateCompact(t *testing.T) {
	ctx := context.Background()
	docs := fixture.Corpus(3, 200)
	data := fixture.Pack(t, docs, nil)
	a, dir, ix := openIndex(t, data)
	ix.Close()
	fixture.Repack(t, data, changed(docs), nil)
	a, err := archive.Open(ctx, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := index.Open(ctx, a, dir, 0); !errors.Is(err, index.ErrStale) {
		t.Fatal(err)
	}
	o := testOpts
	o.Policy.AppendRows = 5
	if _, err := index.Update(ctx, a, dir, o); err != nil {
		t.Fatal(err)
	}
	_, _, fresh := openIndex(t, data)
	check := func(step string) {
		t.Helper()
		ix, err := index.Open(ctx, a, dir, 0)
		if err != nil {
			t.Fatal(err)
		}
		defer ix.Close()
		r := rand.New(rand.NewSource(2))
		for _, q := range randomQueries(r, 150) {
			if x, y := answer(t, ix, false, q...), answer(t, fresh, false, q...); !slices.Equal(x, y) {
				t.Fatalf("%s: %q: %d != %d", step, q, len(x), len(y))
			}
		}
	}
	check("update")
	if _, err := index.Compact(ctx, dir, o, false); err != nil {
		t.Fatal(err)
	}
	check("compact")
	if _, err := index.Compact(ctx, dir, o, true); err != nil {
		t.Fatal(err)
	}
	check("full")
	if err := index.Verify(ctx, a, dir, o); err != nil {
		t.Fatal(err)
	}
}
