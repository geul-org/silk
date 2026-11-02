//ff:func feature=index type=scanner control=sequence
//ff:what TestVerify — a sound index passes; damaged segments and tables, differing rows and differing documents fail
package index

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestVerify(t *testing.T) {
	ctx := context.Background()
	data, dir, a := built(t, 3)
	if err := Verify(ctx, a, dir, testOpts); err != nil {
		t.Fatal(err)
	}
	if err := Verify(ctx, a, dir, Options{}); !errors.Is(err, ErrLimit) {
		t.Fatal(err)
	}
	if err := Verify(ctx, a, t.TempDir(), testOpts); !errors.Is(err, ErrNoIndex) {
		t.Fatal(err)
	}
	ix, _ := Open(ctx, a, dir, 0)
	defer ix.Close()
	if err := cloneDecls(dir, filepath.Join(t.TempDir(), "c"), Manifest{Grammar: "none"}); err == nil {
		t.Fatal("no grammar")
	}
	if err := cloneDecls(dir, filepath.Join(t.TempDir(), "c"), Manifest{Sections: []Section{{Decl: "none"}}}); err == nil {
		t.Fatal("no declaration")
	}
	file := filepath.Join(dir, "f")
	os.WriteFile(file, nil, 0o644)
	if err := cloneDecls(dir, filepath.Join(file, "c"), ix.Manifest); err == nil {
		t.Fatal("under a file")
	}
	other := *ix
	other.Live = Bitmap{}
	if err := ix.same(&other); !errors.Is(err, ErrMismatch) {
		t.Fatal(err)
	}
	var half Bitmap
	half.Put(0, true)
	half.Put(1, true)
	half.Put(2, true)
	lost := *ix
	lost.Live = half
	lost.secs = map[string][]*Sec{"s": {{Files: nil}, {Files: nil}}, "t": {{Files: nil}}}
	if err := ix.same(&lost); !errors.Is(err, ErrMismatch) {
		t.Fatal(err)
	}
	seg0 := filepath.Join(dir, ix.Manifest.Sections[0].Runs[0].File)
	b, _ := os.ReadFile(seg0)
	b[0] ^= 1
	os.WriteFile(seg0, b, 0o644)
	ix3, _ := Open(ctx, a, dir, 0)
	defer ix3.Close()
	if err := ix3.checkFiles(); !errors.Is(err, seg.ErrFormat) {
		t.Fatal(err)
	}
	if err := Verify(ctx, a, dir, testOpts); err == nil {
		t.Fatal("damaged segment")
	}
	b[0] ^= 1
	os.WriteFile(seg0, b, 0o644)
	dp := filepath.Join(dir, ix.Manifest.Docs[0].File)
	d, _ := os.ReadFile(dp)
	d[0] ^= 1
	os.WriteFile(dp, d, 0o644)
	ix4, _ := Open(ctx, a, dir, 0)
	defer ix4.Close()
	if err := ix4.checkFiles(); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	_ = data
}
