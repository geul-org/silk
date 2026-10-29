//ff:func feature=fixture type=builder control=sequence
//ff:what Evolve — write the next generation of a Packed folder: 0x…1d52 rewritten (with 0x…22b7's bytes), 0x…20b2 removed, 0x…9999 added (with 0x…20b2's bytes)
package fixture

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

// Evolve packs a changed copy of Source() onto out's sets (records raw), so
// out holds two generations: an incremental index sees one changed, one
// removed and one new document.
func Evolve(t testing.TB, out string) {
	t.Helper()
	root := t.TempDir()
	if err := os.CopyFS(root, os.DirFS(Source())); err != nil {
		t.Fatal(err)
		return
	}
	base := filepath.Join(root, ".well-known", "geul", "0x00010000")
	b, _ := os.ReadFile(filepath.Join(base, "0x51000000000022b7", "0x00000001.geul"))
	os.WriteFile(filepath.Join(base, "0x5100000000001d52", "0x00000001.geul"), b, 0o644)
	os.Rename(filepath.Join(base, "0x51000000000020b2"), filepath.Join(base, "0x5100000000009999"))
	if _, err := pack.Split(context.Background(), root, filepath.Join(out, ".well-known", "geul"), pack.BuildOptions{Codec: pack.Stored{}}); err != nil {
		t.Fatal(err)
	}
}
