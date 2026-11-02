//ff:func feature=query type=builder control=sequence
//ff:what newCond — a condition opened by a declared section name or a TAG name
package query

import (
	"fmt"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/index"
)

func newCond(ix *index.Index, name string) (Cond, error) {
	if d, ok := ix.Decls[name]; ok {
		return Cond{Name: name, Tags: d.Tags, Decl: &d}, nil
	}
	t, err := decl.Tags(ix.Grammar, name)
	if err != nil {
		return Cond{}, fmt.Errorf("%w: %v", ErrSyntax, err)
	}
	return Cond{Name: name, Tags: t}, nil
}
