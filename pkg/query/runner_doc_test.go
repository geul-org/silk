//ff:func feature=query type=scanner control=sequence
//ff:what TestRunnerDoc — a document frames and counts once; a read failure and a frame failure surface
package query

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/index"
)

func TestRunnerDoc(t *testing.T) {
	a, dir, ix := smallIndex(t)
	r := testRunner(ix)
	if d, err := r.doc(0); err != nil || len(d.Recs) == 0 || len(r.seen) != 1 {
		t.Fatal(err)
	}
	cold, _ := index.Open(context.Background(), a, dir, 0)
	defer cold.Close()
	cold.Manifest.Packs[0].Hash = "gone"
	if _, err := testRunner(cold).doc(0); !errors.Is(err, index.ErrStale) {
		t.Fatal(err)
	}
}
