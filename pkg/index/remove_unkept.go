//ff:func feature=index type=writer control=iteration dimension=1
//ff:what removeUnkept — delete the entries of one index subdirectory that are not in the keep set; the number removed
package index

import (
	"os"
	"path/filepath"
)

// removeUnkept removes sub/<name> files missing from keep.
func removeUnkept(dir, sub string, es []os.DirEntry, keep map[string]bool) int {
	n := 0
	for _, e := range es {
		if keep[sub+"/"+e.Name()] {
			continue
		}
		if os.Remove(filepath.Join(dir, sub, e.Name())) == nil {
			n++
		}
	}
	return n
}
