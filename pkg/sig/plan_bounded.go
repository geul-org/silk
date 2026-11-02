//ff:func feature=sig type=builder control=sequence
//ff:what plan.bounded — the plan as it may spread: past MaxExpand one field spreads, and that is the when field itself when it already picked several plans
package sig

func (p plan) bounded(rows, plans int) plan {
	if rows <= MaxExpand {
		return p
	}
	if plans > 1 {
		return p.limit(0)
	}
	return p.limit(1)
}
