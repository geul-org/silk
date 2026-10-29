//ff:func feature=cli type=command control=sequence
//ff:what indexCompact — run (or with -dry list) the due merges; -idle waits for an idle machine; a pause is exit code 3
package cli

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/geul-org/silk/pkg/idle"
	"github.com/geul-org/silk/pkg/index"
)

// indexCompact runs index.Compact under ctx.
func indexCompact(ctx context.Context, dir string, f indexFlags, stdout, stderr io.Writer) int {
	if f.dry {
		return indexStat(dir, stdout, stderr)
	}
	o := index.Options{Threads: f.threads, Full: f.full}
	if f.idle {
		o.Pause = idlePause(f.threads, stderr, idle.Read)
	}
	r, err := index.Compact(ctx, dir, o)
	printReport(stdout, "compact", r)
	if errors.Is(err, index.ErrPaused) {
		fmt.Fprintln(stderr, "silk: paused")
		return 3
	}
	if err != nil {
		return fail(stderr, err)
	}
	return 0
}
