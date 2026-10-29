//ff:func feature=index type=reader control=iteration dimension=1
//ff:what Open — the catalog's maps with their data packs, after checking the catalog names the archive's current generations
package index

import "github.com/park-jun-woo/silk/pkg/archive"

// Open reads index.json, checks it against a (ErrStale when the archive has
// moved on) and opens every listed map. Listing the sets' documents (State)
// is not needed.
func Open(a *archive.Archive, dir string) (*Index, error) {
	c, err := readCatalog(dir)
	if err != nil {
		return nil, err
	}
	if err := checkSets(a, c); err != nil {
		return nil, err
	}
	ix := &Index{Dir: dir}
	for _, rel := range c.Maps {
		part, err := openPart(a, mapPath(dir, rel))
		if err != nil {
			ix.Close()
			return nil, err
		}
		ix.Parts = append(ix.Parts, part)
	}
	return ix, nil
}
