//ff:func feature=cli type=command control=iteration dimension=1
//ff:what indexBuild — build the index of an archive and print each map (Issuer, pack, documents, records, keys) and totals
package cli

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/park-jun-woo/silk/pkg/archive"
	"github.com/park-jun-woo/silk/pkg/index"
)

// indexBuild runs index.Build and reports.
func indexBuild(arch, dir string, stdout, stderr io.Writer) int {
	t0 := time.Now()
	ctx := context.Background()
	a, err := archive.Open(ctx, arch)
	if err != nil {
		return fail(stderr, err)
	}
	hs, err := index.Build(ctx, a, dir)
	if err != nil {
		return fail(stderr, err)
	}
	var docs, recs uint64
	for _, h := range hs {
		fmt.Fprintf(stdout, "%s %s documents %d records %d keys %d\n", h.Issuer, h.Pack[:12], h.Docs, h.Records, h.Keys)
		docs, recs = docs+h.Docs, recs+h.Records
	}
	fmt.Fprintf(stderr, "silk: %d maps, %d documents, %d records in %s (%v)\n", len(hs), docs, recs, dir, time.Since(t0).Round(time.Millisecond))
	return 0
}
