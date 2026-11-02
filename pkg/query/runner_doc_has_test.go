//ff:func feature=query type=scanner control=sequence
//ff:what TestRunnerDocHas — exact row, candidate row, no row, and a candidate whose document cannot be read
package query

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/index"
)

func TestRunnerDocHas(t *testing.T) {
	a, dir, ix := smallIndex(t)
	sec := ix.Section("entity", 1)
	r := testRunner(ix)
	if ok, err := r.docHas(silkPlanOf(t, ix, "job=physicist"), sec, 0, false); !ok || err != nil {
		t.Fatal(ok, err)
	}
	if ok, err := r.docHas(silkPlanOf(t, ix, "job=physicist"), sec, 1, false); ok || err != nil {
		t.Fatal(ok, err)
	}
	if ok, err := r.docHas(silkPlanOf(t, ix, "type=taxon"), sec, 0, false); ok || err != nil {
		t.Fatal(ok, err)
	}
	if ok, err := r.docHas(silkPlanOf(t, ix, "job=Q98"), sec, 1, false); !ok || err != nil {
		t.Fatal(ok, err)
	}
	if ok, err := r.docHas(silkPlanOf(t, ix, "job=Q98"), sec, 1, true); !ok || err != nil {
		t.Fatal("rows only", ok, err)
	}
	cold, _ := index.Open(context.Background(), a, dir, 0)
	defer cold.Close()
	cold.Manifest.Packs[0].Hash = "gone"
	if _, err := testRunner(cold).docHas(silkPlanOf(t, ix, "job=Q98"), sec, 1, false); !errors.Is(err, index.ErrStale) {
		t.Fatal(err)
	}
}
