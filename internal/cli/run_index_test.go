//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestRunIndex — every subcommand dispatches; missing subcommand, bad flags, missing archive and unknown subcommands are usage errors
package cli

import (
	"strings"
	"testing"
)

func TestRunIndex(t *testing.T) {
	data := packedIndex(t)
	for _, c := range []string{"update", "add", "compact", "verify", "stat"} {
		if code, _, e := call("index " + c + " " + data); code != 0 {
			t.Fatal(c, e)
		}
	}
	if code, o, _ := call("index spec"); code != 0 || !strings.Contains(o, "TRIPLE") {
		t.Fatal(code)
	}
	for _, a := range []string{"", "build", "build -x " + data, "nope " + data} {
		if code := runIndex(strings.Fields(a), &strings.Builder{}, &strings.Builder{}); code != 2 {
			t.Fatal(a, code)
		}
	}
}
