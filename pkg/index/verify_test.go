//ff:func feature=index type=rule control=sequence
//ff:what TestVerify — a built index verifies; no index, an unreadable archive, a wrong bitmap, a damaged document and a damaged segment fail
package index

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestVerify(t *testing.T) {
	ctx := context.Background()
	a, _, dir := built(t)
	r, err := Verify(ctx, a, dir)
	if err != nil || r.Docs != 4 || r.Sections != 7 || r.Segments != 7 || r.Rows == 0 {
		t.Fatal(err, r)
	}
	if _, err := Verify(ctx, a, t.TempDir()); !errors.Is(err, ErrNoIndex) {
		t.Fatal(err)
	}
	m, _ := CurrentManifest(dir)
	live, _ := readBitmap(filepath.Join(dir, m.Live))
	live.Put(1, false)
	os.WriteFile(filepath.Join(dir, m.Live), encodeBitmap(live), 0o644)
	if _, err := Verify(ctx, a, dir); !errors.Is(err, ErrMismatch) {
		t.Fatal(err)
	}
	a2, _, dir2 := built(t)
	segFile := filepath.Join(dir2, firstSeg(dir2))
	b, _ := os.ReadFile(segFile)
	b[0] ^= 1
	os.WriteFile(segFile, b, 0o644)
	if _, err := Verify(ctx, a2, dir2); err == nil {
		t.Fatal("damaged segment")
	}
	a3, _, dir3 := built(t)
	ix, _ := Open(a3, dir3)
	pack, off := ix.paths[0], ix.Docs[0].Off
	ix.Close()
	f, _ := os.OpenFile(pack, os.O_WRONLY, 0)
	f.WriteAt([]byte{0xff, 0xff}, int64(off))
	f.Close()
	if _, err := Verify(ctx, a3, dir3); !errors.Is(err, ErrFormat) {
		t.Fatal("damaged document", err)
	}
	os.Remove(filepath.Join(a3.Sets[0].Dir, "m", a3.Sets[0].Manifest+".gpm"))
	if _, err := Verify(ctx, a3, dir3); err == nil {
		t.Fatal("unreadable archive")
	}
}
