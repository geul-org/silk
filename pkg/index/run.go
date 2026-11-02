//ff:func feature=index type=builder control=sequence
//ff:what run — one locked writer run under the memory limit
package index

import (
	"context"
	"os"
	"path/filepath"
	"runtime/debug"
	"strconv"

	"github.com/geul-org/silk/pkg/archive"
)

func run(ctx context.Context, ar *archive.Archive, dir string, o Options, reset bool) (Report, error) {
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
	if reset {
		m = resetManifest(m)
	}
	u := &updater{ar: ar, dir: dir, o: o, m: m, tmp: filepath.Join(dir, dirTmp, strconv.Itoa(os.Getpid()))}
	defer os.RemoveAll(u.tmp)
	if err := os.MkdirAll(u.tmp, 0o755); err != nil {
		return Report{}, err
	}
	return u.advance(ctx)
}
