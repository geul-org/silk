//ff:func feature=index type=helper control=iteration dimension=1
//ff:what keepCheckpoint — add the unfinished output (.tmp) of every paused merge to the keep set
package index

// keepCheckpoint protects files a resume needs.
func keepCheckpoint(dir string, keep map[string]bool) {
	for _, c := range readCheckpoints(dir) {
		keep[c.Output+".tmp"] = true
	}
}
