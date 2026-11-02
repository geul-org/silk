//ff:func feature=cli type=command control=sequence
//ff:what env.one — one query
package cli

import (
	"fmt"
	"strings"

	"github.com/geul-org/silk/pkg/index"
	"github.com/geul-org/silk/pkg/query"
)

func (e *env) one(ix *index.Index, f *flags, words []string, lead string) error {
	q, err := query.Parse(ix, words)
	if err != nil {
		return err
	}
	q.Mode, q.Scan, q.Index = f.mode(), f.scan, f.name
	r, err := query.Run(e.ctx, ix, q)
	if err != nil {
		return err
	}
	if f.explain {
		fmt.Fprintf(e.out, "%s# %s, %d hits, %d documents opened\n", lead, strings.Join(r.Plans, "; "), len(r.Hits), r.Decoded)
	}
	return e.print(ix, q.Mode, lead, r.Hits)
}
