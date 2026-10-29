//ff:func feature=index type=helper control=iteration dimension=1
//ff:what Index.Section — the section and spec tag of a TAG kind; false when the spec does not index it
package index

import "github.com/park-jun-woo/silk/pkg/spec"

// Section finds the section whose TAG WORD is tag.
func (ix *Index) Section(tag uint16) (int, spec.Tag, bool) {
	for i, s := range ix.M.Sections {
		if s.Tag == tag {
			t, ok := ix.Spec.Find(tag)
			return i, t, ok
		}
	}
	return 0, spec.Tag{}, false
}
