//ff:func feature=cli type=helper control=sequence
//ff:what built — a project built and its opened index
package cli

import (
	"io"
	"testing"

	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/index"
)

func built(t *testing.T) (string, *flags, *index.Index) {
	t.Helper()
	data, ix := project(t)
	f := newFlags("t", io.Discard)
	f.index = ix[1]
	f.mem = "64M"
	e, _ := testEnv("")
	if err := e.write(f, data, true); err != nil {
		t.Fatal(err)
	}
	a, _ := archive.Open(e.ctx, data)
	i, err := index.Open(e.ctx, a, f.dir(data), 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(i.Close)
	return data, f, i
}
