//ff:func feature=cli type=command control=sequence
//ff:what runQuery — `silk query`: parse flags and conditions, open archive and index, print the plan or run it
package cli

import (
	"context"
	"fmt"
	"io"

	"github.com/park-jun-woo/silk/pkg/archive"
	"github.com/park-jun-woo/silk/pkg/index"
	"github.com/park-jun-woo/silk/pkg/query"
	"github.com/park-jun-woo/silk/pkg/scan"
)

// runQuery answers -triple conditions over ARCHIVE's index.
func runQuery(args []string, stdout, stderr io.Writer) int {
	fs, dir := newFlags("query", stderr)
	var ts triples
	fs.Var(&ts, "triple", "condition PRED=OBJ (repeatable)")
	threads := fs.Int("threads", scan.DefaultThreads, "scanning goroutines")
	count := fs.Bool("count", false, "print only the counts")
	explain := fs.Bool("explain", false, "print the plan instead of running it")
	if fs.Parse(args) != nil || fs.NArg() != 1 || len(ts) == 0 {
		fmt.Fprint(stderr, usageText)
		return 2
	}
	q, err := parseQuery(ts)
	if err != nil {
		return fail(stderr, err)
	}
	if *explain {
		printPlan(stdout, query.NewPlan(q, *threads))
		return 0
	}
	ctx := context.Background()
	a, err := archive.Open(ctx, fs.Arg(0))
	if err != nil {
		return fail(stderr, err)
	}
	ix, err := index.Open(a, indexDir(*dir, fs.Arg(0)))
	if err != nil {
		return fail(stderr, err)
	}
	defer ix.Close()
	return execQuery(ctx, ix, q, scan.Options{Threads: *threads}, *count, stdout, stderr)
}
