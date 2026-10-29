//ff:func feature=index type=helper control=sequence
//ff:what Manifest.newName — a fresh relative file name under a subdirectory (the manifest's file counter, 16 hex digits)
package index

import "fmt"

// newName reserves the next file number.
func (m *Manifest) newName(sub, ext string) string {
	m.NextFile++
	return fmt.Sprintf("%s/%016x%s", sub, m.NextFile, ext)
}
