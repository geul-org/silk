//ff:func feature=cli type=command control=sequence
//ff:what runCreate — declare a section; `silkg index update` builds it
package cli

import (
	"fmt"

	"github.com/geul-org/silk/pkg/index"
)

func runCreate(e *env, args []string) error {
	f := newFlags("create", e.err)
	pos, err := f.parseArgs(args)
	if err != nil {
		return err
	}
	if len(pos) < 3 {
		return fmt.Errorf("create: %s", usage)
	}
	ar := pos[len(pos)-1]
	g, err := loadGrammar(e.ctx, f, ar)
	if err != nil {
		return err
	}
	d, err := f.declaration(g, pos[0], pos[1], pos[2:len(pos)-1])
	if err != nil {
		return err
	}
	return index.Declare(f.dir(ar), d, g)
}
