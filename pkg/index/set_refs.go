//ff:func feature=index type=builder control=iteration dimension=1
//ff:what setRefs — the set generations of an archive
package index

import "github.com/geul-org/silk/pkg/archive"

func setRefs(ar *archive.Archive) []SetRef {
	out := make([]SetRef, 0, len(ar.Sets))
	for _, s := range ar.Sets {
		out = append(out, SetRef{Name: s.Name, Generation: s.Generation, Manifest: s.Manifest})
	}
	return out
}
