//ff:func feature=cli type=helper control=sequence
//ff:what TestPrintPlan — numbered stages marked run or skip
package cli

import (
	"strings"
	"testing"

	"github.com/geul-org/silk/pkg/query"
)

func TestPrintPlan(t *testing.T) {
	var b strings.Builder
	printPlan(&b, query.Plan{Steps: []query.Step{{Name: "a", Detail: "x"}, {Name: "b", Active: true, Detail: "y"}}})
	if b.String() != "1. a [skip] x\n2. b [run] y\n" {
		t.Fatal(b.String())
	}
}
