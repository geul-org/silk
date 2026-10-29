//ff:type feature=query type=model
//ff:what CondPlan — how one condition runs: its section, the selection (partition and sort slots) and which fields the keys cover
package query

import "github.com/geul-org/silk/pkg/seg"

// CondPlan is one condition mapped onto the index spec.
type CondPlan struct {
	Cond    Cond
	Section int
	Sel     seg.Sel
	ByPart  bool
	BySort  bool
}
