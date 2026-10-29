//ff:type feature=query type=model
//ff:what Step — one stage of a plan: its name, whether it runs, and what it does
package query

// Step describes a plan stage for explanation.
type Step struct {
	Name   string
	Active bool
	Detail string
}
