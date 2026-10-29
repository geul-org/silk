//ff:func feature=query type=builder control=sequence
//ff:what NewPlan — the stages for a query: signature filter (off until an index carries signatures), TRIPLE map scan, exact check
package query

import "fmt"

// NewPlan lays out how Run answers q with threads scanning goroutines.
func NewPlan(q Query, threads int) Plan {
	return Plan{Query: q, Steps: []Step{
		{Name: "signature filter", Detail: "skipped: the index has no entity signatures yet"},
		{Name: "record map scan", Active: true, Detail: fmt.Sprintf("TRIPLE range (key 0x%04x) of every map, %d condition(s), %d thread(s), Predicate/Object bytes compared in the scan loop", tripleKey, len(q.Triples), threads)},
		{Name: "exact check", Active: true, Detail: "each hit decoded with gogeul core (TRIPLE fields, Edge Id); documents must meet every condition"},
	}}
}
