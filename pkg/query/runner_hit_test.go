//ff:func feature=query type=scanner control=sequence
//ff:what TestRunnerHit — each kind reads its row; a dead document and a failed mask give none
package query

import (
	"testing"

	"github.com/geul-org/silk/pkg/index"
	"github.com/geul-org/silk/pkg/seg"
)

func TestRunnerHit(t *testing.T) {
	_, _, ix := smallIndex(t)
	r := testRunner(ix)
	sel := func(name string, w ...string) (*Plan, *seg.File, []byte) {
		q, err := Parse(ix, w)
		if err != nil {
			t.Fatal(err)
		}
		p, err := makePlan(ix, &q.Conds[0], name)
		if err != nil {
			t.Fatal(err)
		}
		f := ix.Section(name, 0).Files[0]
		_, reg := f.PartAt(0)
		return p, f, reg.Row(0)
	}
	if p, f, row := sel("label", "LABEL"); true {
		if h, ok := r.hit(p, f, row); !ok || h.HasOff {
			t.Fatal(h, ok)
		}
	}
	if p, f, row := sel("label-sem", "LABEL"); true {
		if h, ok := r.hit(p, f, row); !ok || h.QID == 0 {
			t.Fatal(h, ok)
		}
	}
	if p, f, row := sel("triple", "TRIPLE"); true {
		if h, ok := r.hit(p, f, row); !ok || !h.HasOff {
			t.Fatal(h, ok)
		}
	}
	p := silkPlanOf(t, ix, "type=person")
	f := ix.Section("entity", 0).Files[0]
	row := f.All().Row(0)
	if h, ok := r.hit(p, f, row); !ok || h.Key == nil {
		t.Fatal(h, ok)
	}
	live := ix.Live
	ix.Live = index.Bitmap{}
	if _, ok := r.hit(p, f, row); ok {
		t.Fatal("dead")
	}
	ix.Live = live
}
