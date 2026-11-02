//ff:func feature=grammar type=scanner control=sequence
//ff:what TestLoad — names, fields, odd members and languages come from the documents; a directory without specs fails
package grammar

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestLoad(t *testing.T) {
	g := testGrammar(t)
	if !slices.Equal(g.Names["LABEL"], []uint16{0x1004}) || !slices.Equal(g.Names["UTF8-TEXT"], []uint16{0x0840, 0x0841}) || !g.Odd[0x0841] {
		t.Fatal(g.Names["LABEL"], g.Names["UTF8-TEXT"])
	}
	if fs := g.Fields[0x100A]; len(fs) != 4 || fs[3].Name != "rank" || fs[3].Width != 1 {
		t.Fatal(fs)
	}
	if g.Langs["ko"] == 0 {
		t.Fatal("no ko")
	}
	dir := t.TempDir()
	if _, err := Load(context.Background(), dir); !errors.Is(err, ErrNoGrammar) {
		t.Fatal(err)
	}
	if _, err := Load(context.Background(), filepath.Join(dir, "latest.txt", "x")); err == nil {
		t.Fatal("no set")
	}
	os.WriteFile(filepath.Join(dir, "latest.txt"), []byte("junk"), 0o644)
	if _, err := Load(context.Background(), dir); err == nil {
		t.Fatal("bad set")
	}
}
