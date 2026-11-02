//ff:func feature=query type=rule control=sequence
//ff:what isPath — a word resolves as a field path of the condition's TAG
package query

import "github.com/geul-org/silk/pkg/index"

func isPath(ix *index.Index, c *Cond, path string) bool {
	if len(c.Tags) == 0 {
		return false
	}
	_, err := ix.Grammar.Resolve(c.Tags[0], path)
	return err == nil
}
