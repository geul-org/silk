//ff:func feature=cli type=command control=sequence
//ff:what TestRun — no arguments and unknown commands are usage errors; help prints usage; index and query dispatch
package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	var out, errb bytes.Buffer
	if Run(nil, &out, &errb) != 2 || !strings.Contains(errb.String(), "Usage") {
		t.Fatal(errb.String())
	}
	if Run([]string{"nope"}, &out, &errb) != 2 {
		t.Fatal("unknown")
	}
	out.Reset()
	if Run([]string{"help"}, &out, &errb) != 0 || !strings.Contains(out.String(), "silk index build") {
		t.Fatal(out.String())
	}
	if Run([]string{"index"}, &out, &errb) != 2 || Run([]string{"query"}, &out, &errb) != 2 {
		t.Fatal("dispatch")
	}
}
