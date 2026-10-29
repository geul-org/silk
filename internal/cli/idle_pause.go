//ff:func feature=cli type=helper control=sequence
//ff:what idlePause — a pause hook that blocks while the machine is busy (load or memory leave less than 30% free), polling every 2 seconds; says once how idleness is judged
package cli

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/park-jun-woo/silk/pkg/idle"
)

// idlePause samples with read (idle.Read outside tests).
func idlePause(threads int, stderr io.Writer, read func() (idle.Sample, error)) func(context.Context) error {
	r := idle.Rule{Threads: threads}
	if s, err := read(); err == nil {
		fmt.Fprintf(stderr, "silk: idle when (load %.2f + %d threads) / %d CPUs <= 0.70 and memory available >= 30%% (%d of %d MB)\n",
			s.Load1, threads, s.CPUs, s.MemAvail>>20, s.MemTotal>>20)
	} else {
		fmt.Fprintf(stderr, "silk: %v; compacting without waiting\n", err)
	}
	return func(ctx context.Context) error {
		return idle.Wait(ctx, r, 2*time.Second, read)
	}
}
