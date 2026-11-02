//ff:func feature=index type=builder control=iteration dimension=1
//ff:what tierRuns — the runs of the lowest size tier that holds base of them; nil when none does
package index

import (
	"maps"
	"slices"
)

func tierRuns(runs []SegRef, base int) []SegRef {
	if base < 2 {
		return nil
	}
	by := map[int][]SegRef{}
	for _, r := range runs {
		t := tierOf(r.Rows, base)
		by[t] = append(by[t], r)
	}
	keys := slices.Sorted(maps.Keys(by))
	for _, k := range keys {
		if len(by[k]) >= base {
			return by[k]
		}
	}
	return nil
}
