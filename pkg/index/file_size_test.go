//ff:func feature=index type=helper control=sequence
//ff:what fileSize — test helper: a file's size (0 when missing)
package index

import "os"

// fileSize stats path.
func fileSize(path string) int64 {
	fi, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fi.Size()
}
