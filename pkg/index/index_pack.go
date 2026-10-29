//ff:func feature=index type=reader control=sequence
//ff:what Index.pack — a data pack file by its number, opened on first use
package index

import "os"

// pack is safe for concurrent use.
func (ix *Index) pack(i uint32) (*os.File, error) {
	ix.mu.Lock()
	defer ix.mu.Unlock()
	if ix.packs[i] != nil {
		return ix.packs[i], nil
	}
	f, err := os.Open(ix.paths[i])
	if err != nil {
		return nil, err
	}
	ix.packs[i] = f
	return f, nil
}
