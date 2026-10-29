//ff:func feature=scan type=helper control=sequence
//ff:what TestJoin — parts concatenate in order; a chunk error is returned
package scan

import (
	"errors"
	"testing"
)

func TestJoin(t *testing.T) {
	got, err := join([][]Hit{{{Off: 1}}, nil, {{Off: 5}, {Off: 6}}}, make([]error, 3))
	if err != nil || len(got) != 3 || got[2].Off != 6 {
		t.Fatal(got, err)
	}
	bad := errors.New("x")
	if _, err := join([][]Hit{nil, nil}, []error{nil, bad}); err != bad {
		t.Fatal(err)
	}
}
