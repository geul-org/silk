//ff:func feature=query type=decoder control=sequence
//ff:what Query.word — one word: a new condition or a predicate of the last one
package query

import (
	"fmt"

	"github.com/geul-org/silk/pkg/index"
)

func (q *Query) word(ix *index.Index, a string) error {
	path, val, prefix, ok := splitPred(a)
	if !ok {
		c, err := newCond(ix, a)
		q.Conds = append(q.Conds, c)
		return err
	}
	if len(q.Conds) == 0 {
		return fmt.Errorf("%w: %s before any TAG", ErrSyntax, a)
	}
	return addPred(ix, &q.Conds[len(q.Conds)-1], path, val, prefix)
}
