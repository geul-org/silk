//ff:func feature=query type=scanner control=sequence
//ff:what TestRunnerHolds — an entity that holds, none, unknown, and each failure
package query

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/index"
	"github.com/geul-org/silk/pkg/sig"
)

func TestRunnerHolds(t *testing.T) {
	a, dir, ix := smallIndex(t)
	r := testRunner(ix)
	if y, u, err := r.holds(silkPlanOf(t, ix, "job=physicist"), 0); !y || u || err != nil {
		t.Fatal(y, u, err)
	}
	if y, u, err := r.holds(silkPlanOf(t, ix, "sex=male"), 0); y || !u || err != nil {
		t.Fatal(y, u, err)
	}
	if y, u, err := r.holds(silkPlanOf(t, ix, "job=writer"), 0); y || u || err != nil {
		t.Fatal(y, u, err)
	}
	bad := silkPlanOf(t, ix, "job=writer")
	bad.Cond.Silk = []sig.Cond{{Field: "nope"}}
	if _, _, err := r.holds(bad, 0); !errors.Is(err, sig.ErrCond) {
		t.Fatal(err)
	}
	cold, _ := index.Open(context.Background(), a, dir, 0)
	defer cold.Close()
	cold.Manifest.Packs[0].Hash = "gone"
	if _, _, err := testRunner(cold).holds(bad, 0); !errors.Is(err, index.ErrStale) {
		t.Fatal(err)
	}
}
