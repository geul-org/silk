//ff:func feature=cli type=command control=sequence
//ff:what env.stat — print the index's counts
package cli

import (
	"fmt"

	"github.com/geul-org/silk/pkg/index"
)

func (e *env) stat(f *flags, dir string) error {
	st, err := index.ReadStat(f.dir(dir))
	if err != nil {
		return err
	}
	fmt.Fprintf(e.out, "generation %d, document numbers %d, live %d, largest document %d bytes\n", st.Generation, st.Docs, st.Live, st.MaxDoc)
	printSections(e.out, st)
	return nil
}
