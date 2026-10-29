//ff:func feature=index type=reader control=iteration dimension=1
//ff:what pinCurrent — pin the generation CURRENT names, re-reading CURRENT after pinning so a concurrent commit and collection cannot slip between
package index

import (
	"fmt"
	"os"
)

// pinCurrent returns the pinned generation and its pin file; read is
// readCurrent (tests pass a changing one).
func pinCurrent(dir string, read func(string) (uint64, error)) (uint64, string, error) {
	for range 10 {
		g, err := read(dir)
		if err != nil {
			return 0, "", err
		}
		p, err := pinGen(dir, g)
		if err != nil {
			return 0, "", err
		}
		if g2, err := read(dir); err == nil && g2 == g {
			return g, p, nil
		}
		os.Remove(p)
	}
	return 0, "", fmt.Errorf("%w: CURRENT keeps changing", ErrLocked)
}
