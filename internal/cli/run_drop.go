//ff:func feature=cli type=command control=sequence
//ff:what runDrop — remove a section
package cli

import (
	"fmt"

	"github.com/geul-org/silk/pkg/index"
)

func runDrop(e *env, args []string) error {
	f := newFlags("drop", e.err)
	pos, err := f.parseArgs(args)
	if err != nil {
		return err
	}
	if len(pos) != 2 {
		return fmt.Errorf("drop NAME ARCHIVE")
	}
	return index.Drop(f.dir(pos[1]), pos[0])
}
