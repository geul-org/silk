//ff:func feature=index type=writer control=iteration dimension=1
//ff:what collect — remove index files no current or pinned generation names (and no compaction checkpoint keeps); the number removed
package index

import (
	"os"
	"path/filepath"
)

// collect is the garbage collector run after each commit. Old segments stay
// while a reader pins their generation.
func collect(dir string, m Manifest) int {
	keep := map[string]bool{}
	refs(m, keep)
	for g := range pinnedGens(dir) {
		if pm, err := readManifest(dir, g); err == nil {
			refs(pm, keep)
		}
	}
	keepCheckpoint(dir, keep)
	n := 0
	for _, sub := range []string{dirManifest, dirSeg, dirDoc, dirLive, dirSpec} {
		es, _ := os.ReadDir(filepath.Join(dir, sub))
		n += removeUnkept(dir, sub, es, keep)
	}
	return n
}
