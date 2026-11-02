//ff:func feature=query type=scanner control=sequence
//ff:what TestRunnerEntitiesHold — yes, unknown, no, and a bad condition
package query

import (
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/sig"
)

func TestRunnerEntitiesHold(t *testing.T) {
	_, _, ix := smallIndex(t)
	r := testRunner(ix)
	d, _ := r.doc(0)
	def, _ := sig.Parse(testDef)
	if ok, err := r.entitiesHold(def, d, []sig.Cond{{Field: "job", Value: "physicist"}}); !ok || err != nil {
		t.Fatal(ok, err)
	}
	if ok, err := r.entitiesHold(def, d, []sig.Cond{{Field: "job", Value: "writer"}}); ok || err != nil {
		t.Fatal(ok, err)
	}
	if _, err := r.entitiesHold(def, d, []sig.Cond{{Field: "nope", Value: "x"}}); !errors.Is(err, sig.ErrCond) {
		t.Fatal(err)
	}
}
