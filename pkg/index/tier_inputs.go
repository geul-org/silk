//ff:func feature=index type=rule control=iteration dimension=1
//ff:what tierInputs — the runs of the smallest size tier (rows within a power of fanout) holding at least fanout runs; the largest run (the base) never joins
package index

import (
	"cmp"
	"math"
	"slices"
)

// tierInputs groups runs by floor(log_fanout(rows)).
func tierInputs(runs []SegRef, fanout int) []SegRef {
	if len(runs) <= fanout {
		return nil
	}
	rs := append([]SegRef{}, runs...)
	slices.SortStableFunc(rs, func(a, b SegRef) int { return cmp.Compare(b.Rows, a.Rows) })
	tiers := map[int][]SegRef{}
	low := math.MaxInt
	for _, r := range rs[1:] {
		t := int(math.Log(float64(max(r.Rows, 1))) / math.Log(float64(fanout)))
		tiers[t] = append(tiers[t], r)
		if len(tiers[t]) >= fanout {
			low = min(low, t)
		}
	}
	return tiers[low]
}
