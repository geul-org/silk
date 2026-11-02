//ff:func feature=index type=helper control=sequence
//ff:what ingester.anyPending — whether any declaration still needs every document
package index

import "slices"

func (g *ingester) anyPending() bool {
	return slices.ContainsFunc(g.ts, func(t *target) bool { return t.pending })
}
