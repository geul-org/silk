//ff:func feature=grammar type=helper control=sequence
//ff:what junkStore.ReadFile — bytes that are not one document for version documents
package grammar

import "strings"

func (j junkStore) ReadFile(name string) ([]byte, error) {
	b, err := j.Archive.ReadFile(name)
	if strings.HasSuffix(name, "versions.geul") {
		return b, err
	}
	return b[:len(b)-2], err
}
