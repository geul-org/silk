//ff:func feature=index type=builder control=iteration dimension=1
//ff:what newCatalog — the catalog of an archive's current sets and the given map files
package index

import "github.com/park-jun-woo/silk/pkg/archive"

// newCatalog records a's sets in order.
func newCatalog(a *archive.Archive, maps []string) Catalog {
	c := Catalog{Version: catalogVersion, Sets: []CatalogSet{}, Maps: maps}
	for _, s := range a.Sets {
		c.Sets = append(c.Sets, CatalogSet{Name: s.Name, Generation: s.Generation, Manifest: s.Manifest})
	}
	return c
}
