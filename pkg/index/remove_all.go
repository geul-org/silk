//ff:func feature=index type=writer control=iteration dimension=1
//ff:what removeAll — remove files
package index

import "os"

func removeAll(paths []string) {
	for _, p := range paths {
		os.Remove(p)
	}
}
