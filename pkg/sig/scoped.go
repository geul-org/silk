//ff:func feature=sig type=rule control=sequence
//ff:what scoped — a field under `when` answers only on its code; on somevalue or other the entity may still be of that code, so an answer is merely unknown
package sig

import "slices"

func scoped(e *Entity, w *When, f *Field, t Term, v string) (bool, bool) {
	has, some := fieldHolds(e.codes(f), t)
	has = has || e.stated(f, v)
	if w == nil {
		return has, some
	}
	codes := e.codes(w.On)
	lo, _ := w.Entry.span(w.On.Width)
	if slices.Contains(codes, lo) {
		return has, some
	}
	loose := slices.Contains(codes, CodeSomeValue) || slices.Contains(codes, CodeOther)
	return false, loose && (has || some)
}
