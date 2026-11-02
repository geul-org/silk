//ff:func feature=cli type=command control=sequence
//ff:what env.verify — check the index against its archive
package cli

import (
	"fmt"

	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/index"
)

func (e *env) verify(f *flags, dir string) error {
	o, err := f.options()
	if err != nil {
		return err
	}
	a, err := archive.Open(e.ctx, dir)
	if err != nil {
		return err
	}
	if err := index.Verify(e.ctx, a, f.dir(dir), o); err != nil {
		return err
	}
	fmt.Fprintln(e.out, "ok")
	return nil
}
