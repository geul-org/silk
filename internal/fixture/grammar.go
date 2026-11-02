//ff:func feature=fixture type=reader control=sequence
//ff:what Grammar — the test grammar
package fixture

import (
	"testing"

	"github.com/geul-org/silk/pkg/grammar"
)

// Grammar loads the copy of the published grammar set.
func Grammar(t testing.TB) *grammar.Grammar {
	t.Helper()
	return grammarAt(t, GrammarDir())
}
