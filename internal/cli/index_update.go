//ff:func feature=cli type=command control=sequence
//ff:what indexUpdate — bring the index to the archive's current generations (appended segments) and print the report
package cli

import (
	"context"
	"io"

	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/index"
)

// indexUpdate runs index.Update, keeping the index's spec unless -spec.
func indexUpdate(arch, dir string, f indexFlags, stdout, stderr io.Writer) int {
	ctx := context.Background()
	sp, err := loadSpec(f.spec, dir, true)
	if err != nil {
		return fail(stderr, err)
	}
	a, err := archive.Open(ctx, arch)
	if err != nil {
		return fail(stderr, err)
	}
	r, err := index.Update(ctx, a, dir, sp, index.Options{Threads: f.threads, Compact: f.compact})
	printReport(stdout, "update", r)
	if err != nil {
		return fail(stderr, err)
	}
	return 0
}
