//ff:func feature=index type=helper control=iteration dimension=1
//ff:what setRefs — an archive's sets as manifest SetRefs
package index

import "github.com/park-jun-woo/silk/pkg/archive"

// setRefs records a's current generations.
func setRefs(a *archive.Archive) []SetRef {
	out := make([]SetRef, 0, len(a.Sets))
	for _, s := range a.Sets {
		out = append(out, SetRef{Name: s.Name, Generation: s.Generation, Manifest: s.Manifest})
	}
	return out
}
