//ff:func feature=cli type=builder control=sequence
//ff:what flags.idlePause — the pause hook of an idle rule; nil when the user set none
package cli

import (
	"context"
	"time"

	"github.com/geul-org/silk/pkg/idle"
)

func (f *flags) idlePause(ctx context.Context) (func() error, error) {
	if f.idleCPU == 0 && f.idleMem == 0 {
		return nil, nil
	}
	r := idle.Rule{CPUFree: f.idleCPU, MemFree: f.idleMem, Threads: 1}
	if err := r.Check(); err != nil {
		return nil, err
	}
	return func() error { return idle.Wait(ctx, r, time.Second, idle.Read) }, nil
}
