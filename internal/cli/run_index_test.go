//ff:func feature=cli type=scanner control=iteration dimension=1
//ff:what TestRunIndex — dispatch of each subcommand and its usage errors
package cli

import (
	"strings"
	"testing"
)

func TestRunIndex(t *testing.T) {
	e, o := testEnv("")
	data, ix := project(t)
	a := func(s ...string) []string { return append(s, append(ix, data)...) }
	for _, x := range [][]string{a("build", "-mem", "8M"), a("update", "-mem", "8M"), a("compact", "-mem", "8M"), a("verify", "-mem", "8M"), a("stat")} {
		if err := runIndex(e, x); err != nil {
			t.Fatal(x, err)
		}
	}
	if !strings.Contains(o.String(), "generation") {
		t.Fatal(o)
	}
	for _, x := range [][]string{{}, {"build", "-nope"}, {"build"}, a("nope")} {
		if err := runIndex(e, x); err == nil {
			t.Fatal(x)
		}
	}
}
