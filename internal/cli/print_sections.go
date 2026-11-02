//ff:func feature=cli type=writer control=iteration dimension=1
//ff:what printSections — one line per section
package cli

import (
	"fmt"
	"io"

	"github.com/geul-org/silk/pkg/index"
)

func printSections(w io.Writer, st index.Stat) {
	for _, s := range st.Sections {
		state := "built"
		if s.Pending {
			state = "pending"
		}
		fmt.Fprintf(w, "%s\t%s\tbundle=%d\t%s\trows=%d\tbytes=%d\truns=%d\tappends=%d\tstale=%.3f\n", s.Name, s.Kind, s.Bundle, state, s.Rows, s.Bytes, s.Runs, s.Appends, s.Stale)
	}
}
