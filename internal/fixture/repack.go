//ff:func feature=fixture type=builder control=sequence
//ff:what Repack — the next generation of out holding exactly these documents: changed paths reappear, missing ones are deleted
package fixture

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func Repack(t testing.TB, out string, docs []Doc, c pack.Codec) {
	t.Helper()
	if c == nil {
		c = pack.Stored{}
	}
	root := t.TempDir()
	writeTree(t, root, docs)
	if _, err := pack.Split(context.Background(), root, filepath.Join(out, ".well-known", "geul"), pack.BuildOptions{Codec: c}); err != nil {
		t.Fatal(err)
	}
}
