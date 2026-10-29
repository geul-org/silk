//ff:func feature=cli type=command control=sequence
//ff:what runQuery — `silk query`: parse flags and conditions, open archive and index, print the plan or run it
package cli

import (
	"context"
	"fmt"
	"io"

	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/index"
	"github.com/geul-org/silk/pkg/query"
)

// runQuery answers -triple, -match and -label conditions over ARCHIVE's index.
func runQuery(args []string, stdout, stderr io.Writer) int {
	var f queryFlags
	fs := newQueryFlags(stderr, &f)
	if fs.Parse(args) != nil || fs.NArg() != 1 || len(f.triples)+len(f.matches)+len(f.labels) == 0 {
		fmt.Fprint(stderr, usageText)
		return 2
	}
	q, err := parseQuery(f)
	if err != nil {
		return fail(stderr, err)
	}
	ctx := context.Background()
	a, err := archive.Open(ctx, fs.Arg(0))
	if err != nil {
		return fail(stderr, err)
	}
	ix, err := index.Open(a, indexDir(f.dir, fs.Arg(0)))
	if err != nil {
		return fail(stderr, err)
	}
	defer ix.Close()
	p, err := query.NewPlan(ix, q, f.threads)
	if err != nil {
		return fail(stderr, err)
	}
	if f.explain {
		printPlan(stdout, p)
		return 0
	}
	return execQuery(ctx, ix, p, f.threads, f.count, stdout, stderr)
}
