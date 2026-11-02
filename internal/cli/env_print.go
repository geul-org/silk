//ff:func feature=cli type=writer control=iteration dimension=1
//ff:what env.print — hits in the asked unit
package cli

import (
	"fmt"

	"github.com/geul-org/silk/pkg/index"
	"github.com/geul-org/silk/pkg/query"
)

func (e *env) print(ix *index.Index, mode, lead string, hits []query.Hit) error {
	for _, h := range hits {
		line, err := e.line(ix, mode, h)
		if err != nil {
			return err
		}
		fmt.Fprintf(e.out, "%s%s\n", lead, line)
	}
	return nil
}
