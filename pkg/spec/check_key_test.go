//ff:func feature=spec type=rule control=iteration dimension=1
//ff:what TestCheckKey — fixed positions and bounded partition keys pass; widths, positions past the layout, codebook-only sub-records and unbounded partition keys fail
package spec

import (
	"errors"
	"testing"
)

func TestCheckKey(t *testing.T) {
	triple, _ := LayoutOf(KindEdge, 0x0A)
	label, _ := LayoutOf(KindEdge, 0x04)
	cases := []struct {
		l    Layout
		k    Key
		part bool
		ok   bool
	}{
		{triple, Key{}, true, true},
		{triple, Key{Width: 2}, false, false},
		{triple, Key{Path: []int{1}, Width: 1}, true, false},
		{triple, Key{Path: []int{1}, Width: 34}, true, false},
		{triple, Key{Path: []int{1}, Width: 5}, true, false},
		{triple, Key{Path: []int{1, 0, 0}, Width: 6}, false, false},
		{triple, Key{Path: []int{-1}, Width: 6}, false, false},
		{triple, Key{Path: []int{4}, Width: 2}, false, false},
		{triple, Key{Path: []int{3, 0}, Width: 6}, false, false},
		{label, Key{Path: []int{2, 1}, Width: 6}, false, false},
		{triple, Key{Path: []int{3}, Width: 6}, true, false},
		{label, Key{Path: []int{2}, Width: 6}, true, false},
		{triple, Key{Path: []int{3}, Width: 6}, false, true},
		{triple, Key{Path: []int{1}, Width: 6}, true, true},
		{label, Key{Path: []int{2, 0}, Width: 16}, true, true},
	}
	for i, c := range cases {
		err := CheckKey(c.l, c.k, c.part)
		if (err == nil) != c.ok || (err != nil && !errors.Is(err, ErrSpec)) {
			t.Fatalf("case %d: %v", i, err)
		}
	}
}
