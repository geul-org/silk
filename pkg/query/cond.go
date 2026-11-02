//ff:type feature=query type=model
//ff:what Cond — one condition: TAGs and predicates on one record, or silk conditions on one entity
package query

import (
	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/sig"
)

type Cond struct {
	Name  string
	Tags  []uint16
	Preds []Pred
	Silk  []sig.Cond
	Decl  *decl.Decl
}
