//ff:func feature=query type=scanner control=sequence
//ff:what TestRunnerCheckSilk — exact rows answer; a loose row of a document with an exact one goes; a loose one reads statements; a failed read stops
package query

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/index"
)

func TestRunnerCheckSilk(t *testing.T) {
	a, dir, ix := smallIndex(t)
	r := testRunner(ix)
	p := silkPlanOf(t, ix, "job=Q98")
	hs, err := r.checkSilk(p, []Hit{{Doc: 0, Loose: true}, {Doc: 1, Loose: true}, {Doc: 2}, {Doc: 2, Loose: true}})
	if err != nil || len(hs) != 2 || hs[0].Doc != 1 || hs[1].Doc != 2 {
		t.Fatal(hs, err)
	}
	cold, _ := index.Open(context.Background(), a, dir, 0)
	defer cold.Close()
	cold.Manifest.Packs[0].Hash = "gone"
	if _, err := testRunner(cold).checkSilk(p, []Hit{{Doc: 0, Loose: true}}); !errors.Is(err, index.ErrStale) {
		t.Fatal(err)
	}
}
