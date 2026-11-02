//ff:func feature=query type=scanner control=sequence
//ff:what TestRunnerScan — every live document once; dead ones are not read; a failed read stops
package query

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/index"
)

func TestRunnerScan(t *testing.T) {
	a, dir, ix := smallIndex(t)
	r := testRunner(ix)
	q, _ := Parse(ix, []string{"TRIPLE", "predicate=P31"})
	res, err := r.scan(q)
	if err != nil || len(res.Hits) != 2 || res.Decoded != 2 || res.Plans[0] != "full scan" {
		t.Fatal(res, err)
	}
	live := ix.Live
	var one index.Bitmap
	one.Put(1, true)
	ix.Live = one
	r = testRunner(ix)
	if res, err = r.scan(q); err != nil || res.Decoded != 1 {
		t.Fatal(res, err)
	}
	ix.Live = live
	cold, _ := index.Open(context.Background(), a, dir, 0)
	defer cold.Close()
	cold.Manifest.Packs[0].Hash = "gone"
	if _, err := testRunner(cold).scan(q); !errors.Is(err, index.ErrStale) {
		t.Fatal(err)
	}
}
