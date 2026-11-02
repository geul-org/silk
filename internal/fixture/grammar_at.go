//ff:func feature=fixture type=reader control=sequence
//ff:what grammarAt — the grammar of a pack set directory; a failure fails the test
package fixture

import (
	"context"
	"testing"

	"github.com/geul-org/silk/pkg/grammar"
)

// grammarAt loads dir.
func grammarAt(t testing.TB, dir string) *grammar.Grammar {
	t.Helper()
	g, err := grammar.Load(context.Background(), dir)
	if err != nil {
		t.Fatal(err)
		return nil
	}
	return g
}
