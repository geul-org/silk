//ff:func feature=index type=writer control=iteration dimension=1
//ff:what collect — remove index files no current or pinned generation names
package index

import (
	"os"
	"path/filepath"
)

func collect(dir string, m Manifest) {
	keep := map[string]bool{}
	refs(m, keep)
	for g := range pinnedGens(dir) {
		if pm, err := readManifest(dir, g); err == nil {
			refs(pm, keep)
		}
	}
	for _, sub := range []string{dirManifest, dirSeg, dirDoc, dirLive, dirDecl, dirGrammar} {
		es, _ := os.ReadDir(filepath.Join(dir, sub))
		removeUnkept(dir, sub, es, keep)
	}
}
