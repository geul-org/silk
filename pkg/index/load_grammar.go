//ff:func feature=index type=reader control=sequence
//ff:what loadGrammar — the stored grammar of a manifest; empty name is the empty grammar of an index with nothing declared
package index

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/geul-org/silk/pkg/grammar"
)

func loadGrammar(dir, name string) (*grammar.Grammar, error) {
	g := &grammar.Grammar{}
	if name == "" {
		return g, nil
	}
	b, err := os.ReadFile(filepath.Join(dir, dirGrammar, name+".json"))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, g); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFormat, err)
	}
	return g, nil
}
