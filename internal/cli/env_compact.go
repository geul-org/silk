//ff:func feature=cli type=command control=sequence
//ff:what env.compact — merge what is due; with -idle-cpu and -idle-mem wait for an idle machine between steps
package cli

import (
	"fmt"

	"github.com/geul-org/silk/pkg/index"
)

func (e *env) compact(f *flags, dir string) error {
	o, err := f.options()
	if err != nil {
		return err
	}
	if o.Pause, err = f.idlePause(e.ctx); err != nil {
		return err
	}
	r, err := index.Compact(e.ctx, f.dir(dir), o, f.full)
	if err != nil {
		return err
	}
	fmt.Fprintf(e.out, "merges %d\n", r.Merges)
	return nil
}
