//ff:func feature=index type=helper control=iteration dimension=1
//ff:what setOf — the archive set whose current manifest is the given SHA-256 hex
package index

import "github.com/park-jun-woo/silk/pkg/archive"

// setOf finds the set at manifest.
func setOf(a *archive.Archive, manifest string) (archive.Set, bool) {
	for _, s := range a.Sets {
		if s.Manifest == manifest {
			return s, true
		}
	}
	return archive.Set{}, false
}
