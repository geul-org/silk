//ff:func feature=index type=writer control=sequence
//ff:what storeGrammar — fix the grammar of an index at its first declaration
package index

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"path/filepath"

	"github.com/geul-org/silk/pkg/grammar"
)

func storeGrammar(dir string, m *Manifest, g *grammar.Grammar) error {
	if m.Grammar != "" {
		return nil
	}
	b, _ := json.Marshal(g)
	h := sha256.Sum256(b)
	m.Grammar = hex.EncodeToString(h[:])
	return writeAtomic(filepath.Join(dir, dirGrammar, m.Grammar+".json"), b)
}
