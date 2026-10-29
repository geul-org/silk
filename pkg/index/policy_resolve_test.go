//ff:func feature=index type=helper control=sequence
//ff:what TestPolicyResolve — zero fields take the defaults, set ones stay
package index

import "testing"

func TestPolicyResolve(t *testing.T) {
	if p := (Policy{}).resolve(); p != DefaultPolicy() {
		t.Fatal(p)
	}
	q := Policy{AppendRatio: 0.5, MaxRuns: 2, StaleRatio: 0.9, Fanout: 3}
	if q.resolve() != q {
		t.Fatal(q.resolve())
	}
}
