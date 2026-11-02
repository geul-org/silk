//ff:func feature=index type=builder control=sequence
//ff:what Compact — merge what the policy says is due until nothing is; with full, one run per section. Idle waits go in Options.Pause
package index

import (
	"context"
	"os"
	"path/filepath"
	"runtime/debug"
	"strconv"
)

func Compact(ctx context.Context, dir string, o Options, full bool) (Report, error) {
	if o.Limit <= 0 {
		return Report{}, ErrLimit
	}
	unlock, err := lockDir(dir)
	if err != nil {
		return Report{}, err
	}
	defer unlock()
	defer debug.SetMemoryLimit(debug.SetMemoryLimit(o.Limit))
	m, err := currentManifest(dir)
	if err != nil {
		return Report{}, err
	}
	c := &compactor{dir: dir, o: o, m: m, full: full, tmp: filepath.Join(dir, dirTmp, strconv.Itoa(os.Getpid()))}
	defer os.RemoveAll(c.tmp)
	if err := os.MkdirAll(c.tmp, 0o755); err != nil {
		return Report{}, err
	}
	return c.rounds(ctx)
}
