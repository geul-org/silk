//ff:func feature=query type=helper control=iteration dimension=1
//ff:what predAt — the index of the equality predicate on a path; -1 when there is none
package query

import (
	"slices"

	"github.com/geul-org/silk/pkg/grammar"
)

func predAt(c *Cond, path grammar.Path) int {
	for i, q := range c.Preds {
		if slices.Equal(q.Leaf.Path, path) {
			return i
		}
	}
	return -1
}
