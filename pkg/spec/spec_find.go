//ff:func feature=spec type=helper control=iteration dimension=1
//ff:what Spec.Find — the tag declared for a TAG WORD (Kind·Codec bits); false when the spec does not index it
package spec

// Find looks tag up among s.Tags.
func (s Spec) Find(tag uint16) (Tag, bool) {
	for _, t := range s.Tags {
		if TagWord(t.Kind, t.Codec) == tag {
			return t, true
		}
	}
	return Tag{}, false
}
