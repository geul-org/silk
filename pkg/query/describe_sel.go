//ff:func feature=query type=helper control=sequence
//ff:what describeSel — a condition's selection in words: one partition or all, binary search on the sort key or full ranges
package query

// describeSel explains a CondPlan.
func describeSel(cp CondPlan) string {
	part, sort := "every partition", "every row"
	if cp.ByPart {
		part = "one partition"
	}
	if cp.BySort {
		sort = "sort key search"
	}
	return part + ", " + sort
}
