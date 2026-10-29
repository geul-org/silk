//ff:func feature=index type=reader control=iteration dimension=1
//ff:what Index.Close — close segments and data packs and release the generation's pin
package index

import "os"

// Close releases ix.
func (ix *Index) Close() error {
	closeSegs(ix.Segs)
	for _, f := range ix.packs {
		if f != nil {
			f.Close()
		}
	}
	if ix.pin != "" {
		return os.Remove(ix.pin)
	}
	return nil
}
