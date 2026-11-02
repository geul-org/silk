//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what TestIndexLoad — each file the generation names must be there; a stale archive is refused
package index

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/decl"
)

func TestIndexLoad(t *testing.T) {
	_, dir, a := built(t, 1)
	m := current(t, dir)
	load := func(ar *archive.Archive) error {
		ix := &Index{Dir: dir, ar: ar, Decls: map[string]decl.Decl{}, secs: map[string][]*Sec{}}
		defer ix.Close()
		return ix.load(m.Generation)
	}
	if err := load(a); err != nil {
		t.Fatal(err)
	}
	if err := load(&archive.Archive{}); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
	for _, f := range []string{m.Sections[0].Runs[0].File, dirDecl + "/" + m.Sections[0].Decl + ".json", dirGrammar + "/" + m.Grammar + ".json", m.Docs[0].File, m.Live, manifestName(m.Generation)} {
		p := filepath.Join(dir, f)
		b, _ := os.ReadFile(p)
		os.Remove(p)
		if err := load(a); err == nil {
			t.Fatal(f)
		}
		os.WriteFile(p, b, 0o644)
	}
}
