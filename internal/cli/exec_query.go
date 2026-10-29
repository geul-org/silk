//ff:func feature=cli type=command control=iteration dimension=1
//ff:what execQuery — run a query and print each document path with its Edge Ids (unless -count), then counts and time
package cli

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/park-jun-woo/silk/pkg/index"
	"github.com/park-jun-woo/silk/pkg/query"
	"github.com/park-jun-woo/silk/pkg/scan"
)

// execQuery prints `<path> #e1,#e2` lines and a summary on stderr.
func execQuery(ctx context.Context, ix *index.Index, q query.Query, o scan.Options, count bool, stdout, stderr io.Writer) int {
	t0 := time.Now()
	rs, st, err := query.Run(ctx, ix, q, o)
	if err != nil {
		return fail(stderr, err)
	}
	el := time.Since(t0)
	edges := 0
	for _, r := range rs {
		edges += len(r.Edges)
		if !count {
			fmt.Fprintf(stdout, "%s %s\n", r.Path(), edgeList(r.Edges))
		}
	}
	if count {
		fmt.Fprintf(stdout, "%d documents, %d statements\n", len(rs), edges)
	}
	fmt.Fprintf(stderr, "silk: %d documents, %d statements; %d maps, %d TRIPLEs scanned, %d hits; %v\n",
		len(rs), edges, st.Maps, st.Scanned, st.Hits, el.Round(time.Microsecond))
	return 0
}
