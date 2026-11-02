//ff:func feature=sig type=rule control=iteration dimension=1
//ff:what fieldHolds — whether codes meet a term as statements, and whether somevalue could still meet it
package sig

func fieldHolds(codes []uint64, t Term) (bool, bool) {
	has, some := false, false
	for _, c := range codes {
		has = has || allows(t.Truth, c)
		some = some || c == CodeSomeValue && allows(t.Spans, c)
	}
	return has, some
}
