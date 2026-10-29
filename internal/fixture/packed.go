//ff:func feature=fixture type=builder control=sequence
//ff:what Packed — testdata/archive packed into a temporary data folder (one set per Issuer) with the given codec; the folder
package fixture

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

// Packed runs pack.Split over Source() into t.TempDir()/.well-known/geul.
// A nil codec stores records raw (pack.Stored).
func Packed(t testing.TB, c pack.Codec) string {
	t.Helper()
	if c == nil {
		c = pack.Stored{}
	}
	out := t.TempDir()
	if _, err := pack.Split(context.Background(), Source(), filepath.Join(out, ".well-known", "geul"), pack.BuildOptions{Codec: c}); err != nil {
		t.Fatal(err)
	}
	return out
}
