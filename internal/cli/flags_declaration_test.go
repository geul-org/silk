//ff:func feature=cli type=scanner control=sequence
//ff:what TestFlagsDeclaration — every kind, and the ways each refuses
package cli

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestFlagsDeclaration(t *testing.T) {
	g := fixture.Grammar(t)
	dir := t.TempDir()
	def := filepath.Join(dir, "d")
	os.WriteFile(def, []byte(semText), 0o644)
	f := newFlags("t", io.Discard)
	f.def = def
	if d, err := f.declaration(g, "silk", "s", nil); err != nil || d.Kind != "silk" {
		t.Fatal(d, err)
	}
	if d, err := f.declaration(g, "composite", "c", []string{"LABEL"}); err != nil || d.Kind != "composite" {
		t.Fatal(d, err)
	}
	if d, err := f.declaration(g, "text", "x", []string{"LABEL"}); err != nil || d.Kind != "text" {
		t.Fatal(d, err)
	}
	f.key = "object.bytes,subject"
	f.width = "16,6"
	if d, err := f.declaration(g, "index", "i", []string{"LABEL"}); err != nil || len(d.Widths) != 2 {
		t.Fatal(d, err)
	}
	f.width = "x"
	if _, err := f.indexDecl(g, "i", "LABEL"); err == nil {
		t.Fatal("bad width")
	}
	f.key, f.width = "", ""
	if d, err := f.indexDecl(g, "i", "LABEL"); err != nil || len(d.Keys) != 0 {
		t.Fatal(d, err)
	}
	if _, err := f.declaration(g, "nope", "x", []string{"LABEL"}); err == nil {
		t.Fatal("kind")
	}
	if _, err := f.declaration(g, "text", "x", nil); err == nil {
		t.Fatal("no TAG")
	}
	f.def = filepath.Join(dir, "none")
	if _, err := f.declaration(g, "silk", "s", nil); err == nil {
		t.Fatal("no file")
	}
	if _, err := f.declaration(g, "composite", "c", []string{"LABEL"}); err == nil {
		t.Fatal("no file")
	}
}
