//ff:func feature=query type=scanner control=sequence
//ff:what TestRunnerScanDoc — all conditions must hold; a failed read and a bad condition stop
package query

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/index"
)

func TestRunnerScanDoc(t *testing.T) {
	a, dir, ix := smallIndex(t)
	r := testRunner(ix)
	q, _ := Parse(ix, []string{"TRIPLE", "predicate=P31", "TRIPLE", "predicate=P106"})
	if hs, err := r.scanDoc(0, q); err != nil || len(hs) != 2 {
		t.Fatal(hs, err)
	}
	q, _ = Parse(ix, []string{"TRIPLE", "predicate=P31", "TRIPLE", "predicate=P9"})
	if hs, err := r.scanDoc(0, q); err != nil || hs != nil {
		t.Fatal(hs, err)
	}
	cold, _ := index.Open(context.Background(), a, dir, 0)
	defer cold.Close()
	cold.Manifest.Packs[0].Hash = "gone"
	if _, err := testRunner(cold).scanDoc(0, q); !errors.Is(err, index.ErrStale) {
		t.Fatal(err)
	}
}
