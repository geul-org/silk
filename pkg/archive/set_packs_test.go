//ff:func feature=archive type=reader control=sequence
//ff:what TestSetPacks — raw fixture gives one data pack with 4 documents in offset order; manifest, open, state and zstd failures surface
package archive

import (
	"context"
	"encoding/binary"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
	"github.com/park-jun-woo/silk/internal/fixture"
)

func TestSetPacks(t *testing.T) {
	ctx := context.Background()
	a, _ := Open(ctx, fixture.Packed(t, nil))
	s := a.Sets[0]
	ps, err := s.Packs(ctx)
	if err != nil || len(ps) != 1 || len(ps[0].Docs) != 4 || ps[0].Set != s {
		t.Fatal(ps, err)
	}
	p := ps[0]
	fi, err := os.Stat(p.Path)
	if err != nil || uint64(fi.Size()) != p.Size || p.Docs[0].Off >= p.Docs[1].Off || p.Docs[0].Issuer != "0x00010000" {
		t.Fatal(fi, err, p.Docs)
	}
	if _, err := (Set{Dir: t.TempDir(), Manifest: "none"}).Packs(ctx); err == nil {
		t.Fatal("missing manifest must fail")
	}
	bad := Set{Dir: t.TempDir(), Manifest: "bad"}
	os.MkdirAll(filepath.Join(bad.Dir, "m"), 0o755)
	os.WriteFile(filepath.Join(bad.Dir, "m", "bad.gpm"), []byte("bad"), 0o644)
	if _, err := bad.Packs(ctx); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
	mb, _ := os.ReadFile(filepath.Join(s.Dir, "m", s.Manifest+".gpm"))
	renamed := s
	renamed.Manifest = "renamed"
	os.WriteFile(filepath.Join(s.Dir, "m", "renamed.gpm"), mb, 0o644)
	if _, err := renamed.Packs(ctx); err == nil {
		t.Fatal("a manifest whose SHA-256 is not its name must fail")
	}
	b, _ := os.ReadFile(p.Path)
	b[binary.LittleEndian.Uint64(b[24:])+10] ^= 0xFF
	os.Chmod(p.Path, 0o644)
	os.WriteFile(p.Path, b, 0o644)
	if _, err := s.Packs(ctx); !errors.Is(err, pack.ErrCorrupt) {
		t.Fatal(err)
	}
	z, _ := Open(ctx, fixture.Packed(t, &pack.Zstd{}))
	if _, err := z.Sets[0].Packs(ctx); !errors.Is(err, ErrCompressed) {
		t.Fatal(err)
	}
}
