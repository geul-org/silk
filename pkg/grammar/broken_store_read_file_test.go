//ff:func feature=grammar type=helper control=sequence
//ff:what brokenStore.ReadFile — fail for version documents
package grammar

import (
	"errors"
	"strings"
)

func (b brokenStore) ReadFile(name string) ([]byte, error) {
	if strings.HasSuffix(name, "versions.geul") {
		return b.Archive.ReadFile(name)
	}
	return nil, errors.New("broken")
}
