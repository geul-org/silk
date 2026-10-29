//ff:func feature=index type=reader control=iteration dimension=1
//ff:what pinnedGens — generations pinned by live processes; pins of dead processes are removed
package index

import (
	"os"
	"path/filepath"
)

// pinnedGens reads pin/ names "<gen>-<pid>-<time>".
func pinnedGens(dir string) map[uint64]bool {
	out := map[uint64]bool{}
	es, _ := os.ReadDir(filepath.Join(dir, dirPin))
	for _, e := range es {
		g, pid, ok := parsePin(e.Name())
		if !ok || !pidAlive(pid) {
			os.Remove(filepath.Join(dir, dirPin, e.Name()))
			continue
		}
		out[g] = true
	}
	return out
}
