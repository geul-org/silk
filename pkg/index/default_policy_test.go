//ff:func feature=index type=builder control=sequence
//ff:what TestDefaultPolicy — 10% appended, 8 segments, 30% dead, fanout 4
package index

import "testing"

func TestDefaultPolicy(t *testing.T) {
	if p := DefaultPolicy(); p != (Policy{AppendRatio: 0.10, MaxRuns: 8, StaleRatio: 0.30, Fanout: 4}) {
		t.Fatal(p)
	}
}
