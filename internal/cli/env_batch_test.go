//ff:func feature=cli type=scanner control=sequence
//ff:what TestEnvBatch — each line answers with its number; a bad line stops; a too-long line is the scanner's error
package cli

import (
	"strings"
	"testing"
)

func TestEnvBatch(t *testing.T) {
	_, f, ix := built(t)
	e, o := testEnv("TRIPLE predicate=P31 object=Q5\n")
	if err := e.batch(ix, f); err != nil || !strings.HasPrefix(o.String(), "1\t") {
		t.Fatal(err, o)
	}
	e, _ = testEnv("NOPE\n")
	if err := e.batch(ix, f); err == nil {
		t.Fatal("bad line")
	}
	e, _ = testEnv(strings.Repeat("x", 1<<17))
	if err := e.batch(ix, f); err == nil {
		t.Fatal("long line")
	}
}
