//ff:type feature=query type=model
//ff:what Plan — the query, each condition's plan and the stages for explanation
package query

// Plan is what NewPlan makes and Run follows.
type Plan struct {
	Query Query
	Conds []CondPlan
	Steps []Step
}
