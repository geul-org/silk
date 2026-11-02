//ff:func feature=fixture type=scanner control=sequence
//ff:what TestGrammarDir — the grammar copy exists; a damaged copy fails Grammar
package fixture

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGrammarDir(t *testing.T) {
	if _, err := os.Stat(filepath.Join(GrammarDir(), "latest.txt")); err != nil {
		t.Fatal(err)
	}
	if fails(t, func(tb testing.TB) { Grammar(tb) }) {
		t.Fatal("sound grammar")
	}
}
