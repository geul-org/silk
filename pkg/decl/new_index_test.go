//ff:func feature=decl type=scanner control=iteration dimension=1
//ff:what TestNewIndex — defaults, given widths, and each refusal
package decl

import (
	"slices"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestNewIndex(t *testing.T) {
	g := fixture.Grammar(t)
	d, err := NewIndex(g, "l", "LABEL", "object.language", []string{"object.bytes", "subject"}, nil)
	if err != nil || d.PartWidth != 16 || !slices.Equal(d.Widths, []int{32, 6}) {
		t.Fatal(d, err)
	}
	cases := []func() (Decl, error){
		func() (Decl, error) { return NewIndex(g, "x", "NOPE", "", nil, nil) },
		func() (Decl, error) { return NewIndex(g, "x", "LABEL", "nope", nil, nil) },
		func() (Decl, error) { return NewIndex(g, "x", "LABEL", "object.bytes", nil, nil) },
		func() (Decl, error) { return NewIndex(g, "x", "LABEL", "", []string{"nope"}, nil) },
		func() (Decl, error) { return NewIndex(g, "x", "LABEL", "", []string{"subject"}, []int{1, 2}) },
		func() (Decl, error) { return NewIndex(g, "x", "LABEL", "", []string{"subject"}, []int{0}) },
	}
	for i, c := range cases {
		if _, err := c(); err == nil {
			t.Fatal(i)
		}
	}
}
