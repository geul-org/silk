//ff:func feature=cli type=command control=sequence
//ff:what indexBuild — build the index of an archive from scratch and print the report
package cli

import (
	"context"
	"io"

	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/index"
)

// indexBuild runs index.Build with -spec or the default spec.
func indexBuild(arch, dir string, f indexFlags, stdout, stderr io.Writer) int {
	ctx := context.Background()
	sp, err := loadSpec(f.spec, dir, false)
	if err != nil {
		return fail(stderr, err)
	}
	a, err := archive.Open(ctx, arch)
	if err != nil {
		return fail(stderr, err)
	}
	r, err := index.Build(ctx, a, dir, sp, index.Options{Threads: f.threads})
	printReport(stdout, "build", r)
	if err != nil {
		return fail(stderr, err)
	}
	return 0
}
