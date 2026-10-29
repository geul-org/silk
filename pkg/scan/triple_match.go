//ff:func feature=scan type=matcher control=sequence
//ff:what Triple.Match — walk a TRIPLE's field records by their TAG WORDs and compare Predicate and Object bytes
package scan

import "bytes"

// Match skips the TAG (and Length extension) and the Edge Id, compares the
// Predicate, skips the Subject and compares the Object.
func (t Triple) Match(rec []byte) bool {
	p := headLen(rec)
	p, ok := nextField(rec, p)
	pe, ok2 := nextField(rec, p)
	if !ok || !ok2 || (t.Pred != nil && !bytes.Equal(rec[p:pe], t.Pred)) {
		return false
	}
	s, ok := nextField(rec, pe)
	oe, ok2 := nextField(rec, s)
	return ok && ok2 && (t.Obj == nil || bytes.Equal(rec[s:oe], t.Obj))
}
