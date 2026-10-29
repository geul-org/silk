//ff:type feature=query type=model
//ff:what Plan — the query and its stages in order: signature filter, record map scan, exact check
package query

// Plan is what NewPlan makes and Run follows.
type Plan struct {
	Query Query
	Steps []Step
}
