//ff:func feature=cli type=command control=sequence
//ff:what env.write — build or update
package cli

import (
	"fmt"

	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/index"
)

func (e *env) write(f *flags, dir string, build bool) error {
	o, err := f.options()
	if err != nil {
		return err
	}
	a, err := archive.Open(e.ctx, dir)
	if err != nil {
		return err
	}
	do := index.Update
	if build {
		do = index.Build
	}
	r, err := do(e.ctx, a, f.dir(dir), o)
	if err != nil {
		return err
	}
	fmt.Fprintf(e.out, "new documents %d, read %d, rows %d\n", r.NewDocs, r.ReadDocs, r.Rows)
	return nil
}
