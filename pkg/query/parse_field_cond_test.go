//ff:func feature=query type=decoder control=iteration dimension=1
//ff:what TestParseFieldCond — a dotted path and token; no '=', a non-numeric or negative path step, or a bad token fails
package query

import (
	"slices"
	"testing"
)

func TestParseFieldCond(t *testing.T) {
	f, err := parseFieldCond(" 2.0=@ko ")
	if err != nil || !slices.Equal(f.Path, []int{2, 0}) || len(f.Value) != 4 {
		t.Fatal(f, err)
	}
	for _, s := range []string{"2", "a=Q1", "-1=Q1", "1=Z"} {
		if _, err := parseFieldCond(s); err == nil {
			t.Fatal(s)
		}
	}
}
