//ff:func feature=query type=scanner control=sequence
//ff:what TestRunnerScanCond — a silk, a composite with meaning, records, and failures
package query

import (
	"testing"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/sig"
)

func TestRunnerScanCond(t *testing.T) {
	_, _, ix := smallIndex(t)
	r := testRunner(ix)
	d, _ := r.doc(0)
	w := func(s ...string) *Cond {
		q, err := Parse(ix, s)
		if err != nil {
			t.Fatal(err)
		}
		return &q.Conds[0]
	}
	if h, err := r.scanCond(0, d, w("entity", "job=physicist")); err != nil || len(h) != 1 || h[0].HasOff {
		t.Fatal(h, err)
	}
	if h, err := r.scanCond(0, d, w("entity", "job=writer")); err != nil || h != nil {
		t.Fatal(h, err)
	}
	if h, err := r.scanCond(0, d, w("label-sem", `object.bytes="Short"`, "type=person")); err != nil || len(h) != 1 {
		t.Fatal(h, err)
	}
	if h, err := r.scanCond(0, d, w("TRIPLE", "predicate=P31")); err != nil || len(h) != 1 {
		t.Fatal(h, err)
	}
	if _, err := r.scanCond(0, d, &Cond{Decl: &decl.Decl{Silk: "x"}, Silk: []sig.Cond{{}}}); err == nil {
		t.Fatal("bad definition")
	}
}
