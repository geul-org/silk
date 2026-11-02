//ff:func feature=fixture type=builder control=sequence
//ff:what Pack — documents packed into a temporary data folder with the codec (nil stores raw); the folder
package fixture

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func Pack(t testing.TB, docs []Doc, c pack.Codec) string {
	t.Helper()
	out := t.TempDir()
	Repack(t, out, docs, c)
	return out
}
