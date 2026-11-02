//ff:func feature=query type=scanner control=sequence
//ff:what TestRunnerMeaningHolds — no mask passes; a mask reads the statements
package query

import (
	"testing"

	"github.com/geul-org/silk/pkg/sig"
)

func TestRunnerMeaningHolds(t *testing.T) {
	_, _, ix := smallIndex(t)
	r := testRunner(ix)
	d, _ := r.doc(1)
	if ok, err := r.meaningHolds(&Plan{}, d); !ok || err != nil {
		t.Fatal(ok, err)
	}
	def, _ := sig.Parse(testSem)
	p := &Plan{Bits: []sig.Pattern{{}}, Def: def, Cond: &Cond{Silk: []sig.Cond{{Field: "type", Value: "person"}}}}
	if ok, err := r.meaningHolds(p, d); ok || err != nil {
		t.Fatal(ok, err)
	}
}
