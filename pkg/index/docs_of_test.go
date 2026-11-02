//ff:func feature=index type=helper control=iteration dimension=1
//ff:what docsOf — n small documents with a label and an instance-of statement
package index

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/silk/internal/fixture"
)

func docsOf(n int) []fixture.Doc {
	var out []fixture.Doc
	for i := range n {
		s := fixture.Tok(fmt.Sprintf("Q%d", i+1))
		out = append(out, fixture.Doc{ID: uint64(i + 1), Version: 1, Body: []core.Record{
			fixture.Triple(1, fixture.Tok("P31"), s, fixture.Tok("Q5"), 0),
			fixture.Label(2, s, "en", fmt.Sprintf("name %d", i)),
		}})
	}
	return out
}
