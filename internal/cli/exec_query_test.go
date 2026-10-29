//ff:func feature=cli type=command control=sequence
//ff:what TestExecQuery — result lines or counts plus a stderr summary; a failing run is exit 1
package cli

import (
	"bytes"
	"context"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/silk/internal/fixture"
	"github.com/park-jun-woo/silk/pkg/archive"
	"github.com/park-jun-woo/silk/pkg/index"
	"github.com/park-jun-woo/silk/pkg/query"
	"github.com/park-jun-woo/silk/pkg/scan"
)

func TestExecQuery(t *testing.T) {
	ctx := context.Background()
	a, _ := archive.Open(ctx, fixture.Packed(t, nil))
	dir := filepath.Join(t.TempDir(), "idx")
	index.Build(ctx, a, dir)
	ix, _ := index.Open(a, dir)
	defer ix.Close()
	q, _ := parseQuery([]string{"P31=Q5"})
	var out, errb bytes.Buffer
	if execQuery(ctx, ix, q, scan.Options{}, false, &out, &errb) != 0 || strings.Count(out.String(), "\n") != 2 || !strings.Contains(errb.String(), "2 documents, 2 statements") {
		t.Fatal(out.String(), errb.String())
	}
	out.Reset()
	if execQuery(ctx, ix, q, scan.Options{}, true, &out, &errb) != 0 || out.String() != "2 documents, 2 statements\n" {
		t.Fatal(out.String())
	}
	if execQuery(ctx, ix, query.Query{}, scan.Options{}, false, &out, &errb) != 1 {
		t.Fatal("empty query")
	}
}
