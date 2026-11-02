//ff:func feature=index type=scanner control=sequence
//ff:what TestCheckSecs — a sound section passes; a damaged segment stops
package index

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestCheckSecs(t *testing.T) {
	ctx := context.Background()
	_, dir, a := built(t, 2)
	ix, _ := Open(ctx, a, dir, 0)
	defer ix.Close()
	if err := checkSecs(ix.secs["t"]); err != nil {
		t.Fatal(err)
	}
	if err := checkSegs(ix.secs["t"][0].Files); err != nil {
		t.Fatal(err)
	}
	p := filepath.Join(dir, ix.Manifest.Sections[2].Runs[0].File)
	b, _ := os.ReadFile(p)
	b[0] ^= 1
	os.WriteFile(p, b, 0o644)
	bad, _ := Open(ctx, a, dir, 0)
	defer bad.Close()
	if err := checkSecs(bad.secs["t"]); !errors.Is(err, seg.ErrFormat) {
		t.Fatal(err)
	}
	if err := checkSegs(bad.secs["t"][0].Files); !errors.Is(err, seg.ErrFormat) {
		t.Fatal(err)
	}
}
