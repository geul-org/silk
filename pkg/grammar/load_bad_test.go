//ff:func feature=grammar type=scanner control=iteration dimension=1
//ff:what TestLoadBad — a set of no specs, and a set with an unreadable spec document
package grammar

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/pack"
)

func TestLoadBad(t *testing.T) {
	ctx := context.Background()
	src := filepath.Join(repoRoot(), "testdata", "grammar", "0x0000")
	dir := filepath.Join(t.TempDir(), "0x0000")
	if err := os.CopyFS(dir, os.DirFS(src)); err != nil {
		t.Fatal(err)
	}
	es, _ := os.ReadDir(filepath.Join(dir, "o"))
	for _, e := range es {
		if strings.HasSuffix(e.Name(), ".gpk") {
			p := filepath.Join(dir, "o", e.Name())
			os.Chmod(p, 0o644)
			os.Truncate(p, 100)
		}
	}
	if _, err := Load(ctx, dir); err == nil {
		t.Fatal("truncated packs")
	}
	r, err := pack.OpenDir(ctx, src, pack.BuildOptions{})
	if err != nil {
		t.Fatal(err)
	}
	_ = r
	root := t.TempDir()
	body, _ := core.EncodeDocument(nil, nil)
	p := filepath.Join(root, ".well-known", "geul", "0x0000", "0x0123000000000000", "0x00000001.geul")
	os.MkdirAll(filepath.Dir(p), 0o755)
	os.WriteFile(p, body, 0o644)
	wk := filepath.Join(t.TempDir(), "wk")
	if _, err := pack.Split(ctx, root, wk, pack.BuildOptions{Codec: pack.Stored{}}); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(ctx, filepath.Join(wk, "0x0000")); !errors.Is(err, ErrNoGrammar) {
		t.Fatal(err)
	}
}
