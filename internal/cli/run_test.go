//ff:func feature=cli type=command control=sequence
//ff:what TestRun — help and dispatch: no arguments and unknown commands are usage errors (2), help is 0
package cli

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	if code, _, e := call(""); code != 2 || !strings.Contains(e, "Usage") {
		t.Fatal(code)
	}
	if code, o, _ := call("help"); code != 0 || !strings.Contains(o, "silk index build") {
		t.Fatal(code)
	}
	if code, _, _ := call("nope"); code != 2 {
		t.Fatal(code)
	}
	if code, _, _ := call("index"); code != 2 {
		t.Fatal(code)
	}
	if code, _, _ := call("query"); code != 2 {
		t.Fatal(code)
	}
}
