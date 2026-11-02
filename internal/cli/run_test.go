//ff:func feature=cli type=scanner control=sequence
//ff:what TestRun — no command and an unknown command print usage; a failed command exits 1
package cli

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	if c, _, e := silkg(""); c != exitUsage || !strings.Contains(e, "usage") {
		t.Fatal(c, e)
	}
	if c, _, _ := silkg("", "nope"); c != exitUsage {
		t.Fatal(c)
	}
	if c, _, e := silkg("", "drop"); c != exitError || !strings.Contains(e, "silkg:") {
		t.Fatal(c, e)
	}
}
