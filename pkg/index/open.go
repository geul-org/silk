//ff:func feature=index type=reader control=sequence
//ff:what Open — pin the current generation and load it against the archive (stale sets are ErrStale)
package index

import "github.com/geul-org/silk/pkg/archive"

// Open opens dir's current generation for queries. Close releases the pin.
func Open(a *archive.Archive, dir string) (*Index, error) {
	gen, pin, err := pinCurrent(dir, readCurrent)
	if err != nil {
		return nil, err
	}
	ix := &Index{Dir: dir, pin: pin}
	if err := ix.load(a, gen); err != nil {
		ix.Close()
		return nil, err
	}
	return ix, nil
}
