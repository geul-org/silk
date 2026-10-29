//ff:func feature=fixture type=helper control=sequence
//ff:what TestSource — Source points at the testdata archive with its four documents
package fixture

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSource(t *testing.T) {
	ents, err := os.ReadDir(filepath.Join(Source(), ".well-known", "geul", "0x00010000"))
	if err != nil || len(ents) != 4 {
		t.Fatal(len(ents), err)
	}
}
