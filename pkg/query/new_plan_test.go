//ff:func feature=query type=builder control=sequence
//ff:what TestNewPlan — three stages; the signature filter is off, scan and exact check run
package query

import (
	"strings"
	"testing"
)

func TestNewPlan(t *testing.T) {
	p := NewPlan(Query{Triples: []Triple{{Text: "P106=Q169470"}}}, 8)
	if len(p.Steps) != 3 || p.Steps[0].Active || !p.Steps[1].Active || !p.Steps[2].Active || !strings.Contains(p.Steps[1].Detail, "0x100a") {
		t.Fatal(p)
	}
}
