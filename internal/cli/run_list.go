//ff:func feature=cli type=command control=sequence
//ff:what runList — print each section: name, kind, rows, bytes, runs, append rows
package cli

import (
	"fmt"

	"github.com/geul-org/silk/pkg/index"
)

func runList(e *env, args []string) error {
	f := newFlags("list", e.err)
	pos, err := f.parseArgs(args)
	if err != nil {
		return err
	}
	if len(pos) != 1 {
		return fmt.Errorf("list ARCHIVE")
	}
	st, err := index.ReadStat(f.dir(pos[0]))
	if err != nil {
		return err
	}
	printSections(e.out, st)
	return nil
}
