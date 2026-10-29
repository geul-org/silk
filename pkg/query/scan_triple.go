//ff:func feature=query type=scanner control=iteration dimension=1
//ff:what scanTriple — one condition over every map: the documents holding a matching TRIPLE and their Edge Ids
package query

import (
	"context"

	"github.com/park-jun-woo/silk/pkg/index"
	"github.com/park-jun-woo/silk/pkg/scan"
)

// scanTriple runs scanPart for each part, stopping when ctx is done.
func scanTriple(ctx context.Context, ix *index.Index, t Triple, o scan.Options, st *Stats) (map[docKey][]uint64, error) {
	out := map[docKey][]uint64{}
	for _, p := range ix.Parts {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		if err := scanPart(p, t, o, st, out); err != nil {
			return nil, err
		}
	}
	return out, nil
}
