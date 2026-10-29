//ff:func feature=index type=helper control=sequence
//ff:what Policy.resolve — zero fields replaced by DefaultPolicy's
package index

// resolve fills p's zero values.
func (p Policy) resolve() Policy {
	d := DefaultPolicy()
	if p.AppendRatio <= 0 {
		p.AppendRatio = d.AppendRatio
	}
	if p.MaxRuns <= 0 {
		p.MaxRuns = d.MaxRuns
	}
	if p.StaleRatio <= 0 {
		p.StaleRatio = d.StaleRatio
	}
	if p.Fanout < 2 {
		p.Fanout = d.Fanout
	}
	return p
}
