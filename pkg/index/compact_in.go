//ff:func feature=index type=writer control=iteration dimension=1
//ff:what compact — merge rounds under a held lock: resumed checkpoints first, then planned tasks, each round after the pause hook lets it, one committed generation per round
package index

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/park-jun-woo/silk/pkg/spec"
)

// compact edits *m generation by generation.
func compact(ctx context.Context, dir string, m *Manifest, o Options) (Report, error) {
	sp, err := spec.Load(filepath.Join(dir, specName(m.Spec)))
	if err != nil {
		return Report{}, err
	}
	live, err := readBitmap(filepath.Join(dir, m.Live))
	if err != nil {
		return Report{}, err
	}
	rep := Report{Generation: m.Generation, Docs: live.Count()}
	tasks, cps := resumeTasks(dir, m, readCheckpoints(dir))
	for {
		if len(tasks) == 0 {
			tasks = planNamed(m, sp, o)
		}
		if len(tasks) == 0 {
			return rep, nil
		}
		if err := waitIdle(ctx, o); err != nil {
			rep.Paused = true
			return rep, errors.Join(ErrPaused, err)
		}
		outs := runRound(ctx, dir, tasks, cps, live, o)
		if err := finishRound(dir, m, outs, &rep); err != nil {
			return rep, err
		}
		tasks, cps = nil, nil
	}
}
