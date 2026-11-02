//ff:func feature=cli type=command control=sequence
//ff:what runQuery — parse conditions, answer, print
package cli

import (
	"fmt"

	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/index"
)

func runQuery(e *env, args []string) error {
	f := newFlags("query", e.err)
	pos, err := f.parseArgs(args)
	if err != nil {
		return err
	}
	if len(pos) < 1 {
		return fmt.Errorf("query COND… ARCHIVE")
	}
	dir := pos[len(pos)-1]
	cache, err := parseSize(f.cache)
	if err != nil {
		return err
	}
	a, err := archive.Open(e.ctx, dir)
	if err != nil {
		return err
	}
	ix, err := index.Open(e.ctx, a, f.dir(dir), cache)
	if err != nil {
		return err
	}
	defer ix.Close()
	if f.batch {
		return e.batch(ix, f)
	}
	return e.one(ix, f, pos[:len(pos)-1], "")
}
