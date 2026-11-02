//ff:func feature=sig type=builder control=iteration dimension=1
//ff:what Entity.codes — the sorted distinct codes of a field; a field with no statement is CodeNone
package sig

import "slices"

func (e *Entity) codes(f *Field) []uint64 {
	os := e.Objects[string(f.Pred)]
	if len(os) == 0 {
		return []uint64{CodeNone}
	}
	var out []uint64
	for _, o := range os {
		out = append(out, f.codesOf(o)...)
	}
	slices.Sort(out)
	return slices.Compact(out)
}
