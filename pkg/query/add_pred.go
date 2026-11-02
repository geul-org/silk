//ff:func feature=query type=builder control=sequence
//ff:what addPred — one predicate of the open condition: a silk field when the condition names one, else a field path
package query

import (
	"fmt"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/index"
	"github.com/geul-org/silk/pkg/sig"
)

func addPred(ix *index.Index, c *Cond, path, val string, prefix bool) error {
	if c.Decl != nil && c.Decl.Silk != "" && !isPath(ix, c, path) {
		c.Silk = append(c.Silk, sig.Cond{Field: path, Value: val})
		return nil
	}
	if c.Decl != nil && c.Decl.Kind == decl.KindSilk {
		return fmt.Errorf("%w: %s on a silk", ErrSyntax, path)
	}
	l, err := ix.Grammar.Resolve(c.Tags[0], path)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrSyntax, err)
	}
	vals, err := ix.Grammar.Encode(l, val)
	if err != nil {
		return err
	}
	c.Preds = append(c.Preds, Pred{Text: path + "=" + val, Leaf: l, Vals: vals, Prefix: prefix})
	return nil
}
