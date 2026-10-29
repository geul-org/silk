//ff:func feature=cli type=command control=sequence
//ff:what TestExecQuery — result lines or counts and a summary; a cancelled context fails
package cli

import (
	"context"
	"io"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/index"
	"github.com/geul-org/silk/pkg/query"
)

func TestExecQuery(t *testing.T) {
	data := packedIndex(t)
	ctx := context.Background()
	a, _ := archive.Open(ctx, data)
	ix, _ := index.Open(a, filepath.Join(data, ".silk"))
	defer ix.Close()
	c, _ := query.ParseTriple("P31=Q5")
	p, _ := query.NewPlan(ix, query.Query{Conds: []query.Cond{c}}, 1)
	var o, e strings.Builder
	if code := execQuery(ctx, ix, p, 1, false, &o, &e); code != 0 || !strings.Contains(o.String(), "#15") || !strings.Contains(e.String(), "2 verified") {
		t.Fatal(code, o.String())
	}
	o.Reset()
	if code := execQuery(ctx, ix, p, 1, true, &o, io.Discard); code != 0 || o.String() != "2 documents, 2 statements\n" {
		t.Fatal(code, o.String())
	}
	cctx, cancel := context.WithCancel(ctx)
	cancel()
	if code := execQuery(cctx, ix, p, 1, false, io.Discard, io.Discard); code != 1 {
		t.Fatal(code)
	}
}
