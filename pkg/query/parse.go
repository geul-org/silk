//ff:func feature=query type=decoder control=iteration dimension=1
//ff:what Parse — turn words into a Query: a word without `=` opens a condition (TAG, silk or section name), `path=value` and `path^=value` fill it
package query

import (
	"fmt"

	"github.com/geul-org/silk/pkg/index"
)

func Parse(ix *index.Index, args []string) (Query, error) {
	q := Query{Mode: ModeDocs}
	for _, a := range args {
		if err := q.word(ix, a); err != nil {
			return q, err
		}
	}
	if len(q.Conds) == 0 {
		return q, fmt.Errorf("%w: no condition", ErrSyntax)
	}
	return q, nil
}
