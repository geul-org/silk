//ff:func feature=query type=scanner control=iteration dimension=1
//ff:what TestRunRandom — 300 random queries equal the full scan, over raw and zstd packs, and after update and compaction
package query

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
	"github.com/geul-org/silk/internal/fixture"
)

func TestRunRandom(t *testing.T) {
	docs := fixture.Corpus(7, 300)
	_, _, ix := openIndex(t, fixture.Pack(t, docs, nil))
	z, err := pack.ParseCodec("zstd")
	if err != nil {
		t.Fatal(err)
	}
	_, _, iz := openIndex(t, fixture.Pack(t, docs, z))
	r := rand.New(rand.NewSource(1))
	for _, q := range randomQueries(r, 300) {
		a, s, b := answer(t, ix, false, q...), answer(t, ix, true, q...), answer(t, iz, false, q...)
		if !slices.Equal(a, s) || !slices.Equal(a, b) {
			t.Fatalf("%q: index %d, scan %d, zstd %d\n%v\n%v", q, len(a), len(s), len(b), a, s)
		}
	}
}
