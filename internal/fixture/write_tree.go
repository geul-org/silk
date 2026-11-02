//ff:func feature=fixture type=writer control=iteration dimension=1
//ff:what writeTree — documents as a directory archive under root
package fixture

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func writeTree(t testing.TB, root string, docs []Doc) {
	t.Helper()
	for _, d := range docs {
		b, err := core.EncodeDocument(nil, d.Body)
		if err != nil {
			t.Fatal(err)
		}
		p := filepath.Join(root, ".well-known", "geul", fmt.Sprintf("%s/0x%016x/0x%08x.geul", issuer, d.ID, d.Version))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, b, 0o644); err != nil {
			t.Fatal(err)
		}
	}
}
