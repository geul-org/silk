//ff:func feature=index type=scanner control=sequence
//ff:what TestOpenIndex — reading a built index: paths, documents, stale and damaged states
package index

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/silk/pkg/archive"
)

func TestOpenIndex(t *testing.T) {
	ctx := context.Background()
	_, dir, a := built(t, 2)
	ix, err := Open(ctx, a, dir, 1<<20)
	if err != nil {
		t.Fatal(err)
	}
	defer ix.Close()
	if ix.Section("s", 1) == nil || ix.Section("s", 2) != nil || ix.Section("none", 0) != nil {
		t.Fatal("sections")
	}
	if p := ix.Path(0); !strings.HasPrefix(p, "0x00010000/") {
		t.Fatal(p)
	}
	b, err := ix.Bytes(ctx, 0)
	if err != nil || len(b) == 0 {
		t.Fatal(err)
	}
	if c, ok := ix.cache.get(0); !ok || !bytes.Equal(c, b) {
		t.Fatal("cache")
	}
	if _, err := ix.Bytes(ctx, 0); err != nil {
		t.Fatal(err)
	}
	ix.Manifest.Packs[0].Hash = "gone"
	if _, err := ix.Bytes(ctx, 1); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
	if _, err := Open(ctx, a, t.TempDir(), 0); !errors.Is(err, ErrNoIndex) {
		t.Fatal(err)
	}
	m := current(t, dir)
	broke := func(p string, want error) {
		t.Helper()
		full := filepath.Join(dir, p)
		b, _ := os.ReadFile(full)
		os.Remove(full)
		if _, err := Open(ctx, a, dir, 0); err == nil {
			t.Fatal("removed", p)
		}
		os.WriteFile(full, b, 0o644)
	}
	broke(m.Live, nil)
	broke(m.Docs[0].File, nil)
	broke(dirGrammar+"/"+m.Grammar+".json", nil)
	broke(m.Sections[0].Runs[0].File, nil)
	broke(dirDecl+"/"+m.Sections[0].Decl+".json", nil)
	broke(manifestName(m.Generation), nil)
	if ix2, err := Open(ctx, a, dir, 0); err != nil {
		t.Fatal(err)
	} else {
		ix2.Close()
	}
	(&archive.Archive{}).Set("x")
	if _, err := Open(ctx, &archive.Archive{}, dir, 0); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
}
