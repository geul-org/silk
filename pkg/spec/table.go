//ff:func feature=spec type=builder control=iteration dimension=1
//ff:what Spec.Table — a TAG WORD (Kind·Codec bits) to tag index lookup table, -1 for kinds not indexed
package spec

// Table has 1<<13 entries, indexed by TagWord(kind, codec).
func (s Spec) Table() []int {
	t := make([]int, 1<<13)
	for i := range t {
		t[i] = -1
	}
	for i, g := range s.Tags {
		t[TagWord(g.Kind, g.Codec)] = i
	}
	return t
}
