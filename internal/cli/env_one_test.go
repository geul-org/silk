//ff:func feature=cli type=scanner control=sequence
//ff:what TestEnvOne — parse and run errors stop; -explain prints the plan
package cli

import (
	"strings"
	"testing"
)

func TestEnvOne(t *testing.T) {
	_, f, ix := built(t)
	e, o := testEnv("")
	f.explain = true
	if err := e.one(ix, f, []string{"TRIPLE", "predicate=P31", "object=Q5"}, ""); err != nil || !strings.Contains(o.String(), "# triple") {
		t.Fatal(err, o)
	}
	if err := e.one(ix, f, []string{"NOPE"}, ""); err == nil {
		t.Fatal("bad word")
	}
	if err := e.one(ix, f, []string{"DESCRIPTION", "subject=Q1"}, ""); err == nil {
		t.Fatal("no plan")
	}
}
