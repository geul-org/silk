//ff:func feature=grammar type=rule control=sequence
//ff:what SlotExact — equal slots mean equal values: the value fits its slot and an empty rest field is not taken for a missing one
package grammar

func SlotExact(v Val, w int) bool {
	if v.Rest {
		return len(v.B) > 0 && len(v.B) <= w-1
	}
	return len(v.B) <= w
}
