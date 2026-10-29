//ff:func feature=cli type=helper control=sequence
//ff:what TestPrintPlan — numbered stages marked run or skip
package cli

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/silk/pkg/query"
)

func TestPrintPlan(t *testing.T) {
	var out bytes.Buffer
	printPlan(&out, query.Plan{Steps: []query.Step{{Name: "a", Detail: "x"}, {Name: "b", Active: true, Detail: "y"}}})
	if out.String() != "1. a [skip] x\n2. b [run] y\n" {
		t.Fatal(out.String())
	}
}
