//ff:func feature=query type=scanner control=sequence
//ff:what TestRunnerCheckDoc — offsets pick records, rows without them scan; a mask stops the document; a read failure surfaces
package query

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/index"
	"github.com/geul-org/silk/pkg/sig"
)

func TestRunnerCheckDoc(t *testing.T) {
	a, dir, ix := smallIndex(t)
	r := testRunner(ix)
	q, _ := Parse(ix, []string{"LABEL", `object.bytes="Short"`})
	p := &Plan{Cond: &q.Conds[0], Decl: ix.Decls["label"]}
	if hs, err := r.checkDoc(p, []Hit{{Doc: 0}}); err != nil || len(hs) != 1 || hs[0].QID == 0 {
		t.Fatal(hs, err)
	}
	if hs, err := r.checkDoc(p, []Hit{{Doc: 0, HasOff: true, Off: 1}}); err != nil || len(hs) != 0 {
		t.Fatal(hs, err)
	}
	def, _ := sig.Parse(testSem)
	m := &Plan{Cond: &Cond{Silk: []sig.Cond{{Field: "type", Value: "taxon"}}}, Decl: ix.Decls["label"], Bits: []sig.Pattern{{}}, Def: def}
	if hs, err := r.checkDoc(m, []Hit{{Doc: 0}}); err != nil || len(hs) != 0 {
		t.Fatal(hs, err)
	}
	m.Cond.Silk = []sig.Cond{{Field: "nope", Value: "x"}}
	if _, err := r.checkDoc(m, []Hit{{Doc: 0}}); !errors.Is(err, sig.ErrCond) {
		t.Fatal(err)
	}
	cold, _ := index.Open(context.Background(), a, dir, 0)
	defer cold.Close()
	cold.Manifest.Packs[0].Hash = "gone"
	if _, err := testRunner(cold).checkDoc(p, []Hit{{Doc: 0}}); !errors.Is(err, index.ErrStale) {
		t.Fatal(err)
	}
	if _, err := testRunner(cold).check(p, []Hit{{Doc: 0}}); !errors.Is(err, index.ErrStale) {
		t.Fatal(err)
	}
}
