//ff:func feature=cli type=scanner control=sequence
//ff:what TestFlagsOptions — the limit is required and parsed; the policy comes from the flags
package cli

import (
	"errors"
	"io"
	"testing"

	"github.com/geul-org/silk/pkg/index"
)

func TestFlagsOptions(t *testing.T) {
	f := newFlags("t", io.Discard)
	if _, err := f.options(); !errors.Is(err, index.ErrLimit) {
		t.Fatal(err)
	}
	f.mem = "x"
	if _, err := f.options(); err == nil {
		t.Fatal("bad size")
	}
	f.mem, f.appendRows, f.runs, f.stale = "2G", 5, 3, 0.5
	o, err := f.options()
	if err != nil || o.Limit != 1<<31 || o.Policy != (index.Policy{AppendRows: 5, RunsPerTier: 3, StaleRatio: 0.5}) {
		t.Fatal(o, err)
	}
}
