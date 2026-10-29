//ff:func feature=cli type=command control=iteration dimension=1
//ff:what execQuery — run a plan and print each document path with its Edge Ids (unless -count), then counts and time
package cli

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/geul-org/silk/pkg/index"
	"github.com/geul-org/silk/pkg/query"
)

// execQuery prints `<path> #e1,#e2` lines and a summary on stderr.
func execQuery(ctx context.Context, ix *index.Index, p query.Plan, threads int, count bool, stdout, stderr io.Writer) int {
	t0 := time.Now()
	rs, st, err := query.Run(ctx, ix, p, threads)
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
	fmt.Fprintf(stderr, "silk: %d documents, %d statements; %d segments, %d rows examined, %d dead skipped, %d candidates, %d verified; %v\n",
		len(rs), edges, st.Segments, st.Rows, st.Dead, st.Candidates, st.Verified, el.Round(time.Microsecond))
	return 0
}
