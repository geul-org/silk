//ff:func feature=cli type=command control=sequence
//ff:what indexVerify — check the index against the archive and print what was checked
package cli

import (
	"context"
	"fmt"
	"io"

	"github.com/park-jun-woo/silk/pkg/archive"
	"github.com/park-jun-woo/silk/pkg/index"
)

// indexVerify runs index.Verify.
func indexVerify(arch, dir string, stdout, stderr io.Writer) int {
	ctx := context.Background()
	a, err := archive.Open(ctx, arch)
	if err != nil {
		return fail(stderr, err)
	}
	r, err := index.Verify(ctx, a, dir)
	if err != nil {
		return fail(stderr, err)
	}
	fmt.Fprintf(stdout, "silk: generation %d ok: %d sections, %d segments, %d documents, %d rows\n", r.Generation, r.Sections, r.Segments, r.Docs, r.Rows)
	return 0
}
