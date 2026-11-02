//ff:func feature=cli type=scanner control=iteration dimension=1
//ff:what TestRunList — sections print; bad arguments and a missing index fail
package cli

import (
	"strings"
	"testing"
)

func TestRunList(t *testing.T) {
	e, o := testEnv("")
	data, ix := project(t)
	if err := runList(e, append(ix, data)); err != nil || !strings.Contains(o.String(), "entity") {
		t.Fatal(err, o)
	}
	for _, a := range [][]string{{"-nope"}, {}, {"-index", t.TempDir(), data}} {
		if err := runList(e, a); err == nil {
			t.Fatal(a)
		}
	}
}
