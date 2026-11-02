//ff:func feature=archive type=scanner control=sequence
//ff:what TestOpenSet — manifest problems stop a set
package archive

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestOpenSet(t *testing.T) {
	_, a := testArchive(t)
	s := a.Sets[0]
	si := pack.SetInfo{Name: s.Name, Dir: s.Dir, Manifest: s.Manifest, Generation: s.Generation}
	if g, err := openSet(context.Background(), si); err != nil || g.Manifest != s.Manifest {
		t.Fatal(err)
	}
	p := filepath.Join(s.Dir, "m", s.Manifest+".gpm")
	b, _ := os.ReadFile(p)
	os.Chmod(p, 0o644)
	os.WriteFile(p, b[:10], 0o644)
	if _, err := openSet(context.Background(), si); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
	b2 := slices.Clone(b)
	b2[len(b2)-1] ^= 1
	os.WriteFile(p, b2, 0o644)
	if _, err := openSet(context.Background(), si); err == nil {
		t.Fatal("hash mismatch")
	}
	os.Remove(p)
	if _, err := openSet(context.Background(), si); err == nil {
		t.Fatal("no manifest")
	}
}
