//ff:func feature=cli type=command control=sequence
//ff:what compactCmd — `silk index compact` with Ctrl-C turned into a checkpoint (the context ends, merges pause at the next partition)
package cli

import (
	"context"
	"io"
	"os"
	"os/signal"
)

// compactCmd runs indexCompact under an interrupt-aware context.
func compactCmd(dir string, f indexFlags, stdout, stderr io.Writer) int {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	return indexCompact(ctx, dir, f, stdout, stderr)
}
