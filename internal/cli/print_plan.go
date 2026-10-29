//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what printPlan — the plan's stages, one per line, marked run or skip
package cli

import (
	"fmt"
	"io"

	"github.com/geul-org/silk/pkg/query"
)

// printPlan writes `N. name [run|skip] detail`.
func printPlan(w io.Writer, p query.Plan) {
	for i, s := range p.Steps {
		mark := "skip"
		if s.Active {
			mark = "run"
		}
		fmt.Fprintf(w, "%d. %s [%s] %s\n", i+1, s.Name, mark, s.Detail)
	}
}
