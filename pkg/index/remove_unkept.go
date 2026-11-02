//ff:func feature=index type=writer control=iteration dimension=1
//ff:what removeUnkept — delete the entries of one subdirectory that are not in keep
package index

import (
	"os"
	"path/filepath"
)

func removeUnkept(dir, sub string, es []os.DirEntry, keep map[string]bool) {
	for _, e := range es {
		if !keep[sub+"/"+e.Name()] {
			os.Remove(filepath.Join(dir, sub, e.Name()))
		}
	}
}
