//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what TestIndexBytes — a read that fails in the pack surfaces
package index

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestIndexBytes(t *testing.T) {
	ctx := context.Background()
	_, dir, a := built(t, 1)
	ix, _ := Open(ctx, a, dir, 0)
	defer ix.Close()
	es, _ := os.ReadDir(filepath.Join(a.Sets[0].Dir, "o"))
	for _, e := range es {
		os.Rename(filepath.Join(a.Sets[0].Dir, "o", e.Name()), filepath.Join(a.Sets[0].Dir, "o", e.Name()+"x"))
	}
	if _, err := ix.Bytes(ctx, 0); err == nil {
		t.Fatal("no pack")
	}
}
