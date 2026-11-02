//ff:func feature=query type=scanner control=sequence
//ff:what TestRunnerFileRows — rows of a segment under a plan; dead documents drop
package query

import (
	"testing"

	"github.com/geul-org/silk/pkg/index"
)

func TestRunnerFileRows(t *testing.T) {
	_, _, ix := smallIndex(t)
	r := testRunner(ix)
	f := ix.Section("entity", 0).Files[0]
	p := silkPlanOf(t, ix, "job=physicist")
	if hs := r.fileRows(p, f); len(hs) == 0 {
		t.Fatal(hs)
	}
	live := ix.Live
	ix.Live = index.Bitmap{}
	if hs := r.fileRows(p, f); len(hs) != 0 {
		t.Fatal(hs)
	}
	ix.Live = live
}
