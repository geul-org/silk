//ff:func feature=query type=helper control=iteration dimension=1
//ff:what suggest — the create index command that would answer
package query

import (
	"fmt"
	"strings"
)

func suggest(c *Cond) string {
	keys := make([]string, 0, len(c.Preds))
	for _, p := range c.Preds {
		keys = append(keys, p.Leaf.Path.String())
	}
	return fmt.Sprintf("silkg create index NAME %s -key %s ARCHIVE", c.Name, strings.Join(keys, ","))
}
