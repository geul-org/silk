//ff:func feature=query type=builder control=sequence
//ff:what silkPlanOf — the plan of `entity field=value`
package query

import (
	"testing"

	"github.com/geul-org/silk/pkg/index"
)

func silkPlanOf(t *testing.T, ix *index.Index, words ...string) *Plan {
	t.Helper()
	q, err := Parse(ix, append([]string{"entity"}, words...))
	if err != nil {
		t.Fatal(err)
	}
	p, err := makePlan(ix, &q.Conds[0], "")
	if err != nil {
		t.Fatal(err)
	}
	return p
}
