//ff:func feature=index type=scanner control=sequence
//ff:what TestDeclare — a name is taken once, a bad declaration or a locked or broken index fails; Drop removes and refuses the unknown
package index

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/decl"
)

func TestDeclare(t *testing.T) {
	_, dir, _ := setup(t, 1)
	g := fixture.Grammar(t)
	s, _ := decl.NewSilk("s", silkDef)
	if err := Declare(dir, s, g); !errors.Is(err, ErrExists) {
		t.Fatal(err)
	}
	if err := Declare(dir, decl.Decl{Name: "x", Kind: "nope"}, g); !errors.Is(err, decl.ErrDecl) {
		t.Fatal(err)
	}
	un, _ := lockDir(dir)
	if err := Declare(dir, s, g); !errors.Is(err, ErrLocked) {
		t.Fatal(err)
	}
	if err := Drop(dir, "s"); !errors.Is(err, ErrLocked) {
		t.Fatal(err)
	}
	un()
	if err := Drop(dir, "nope"); !errors.Is(err, ErrUnknown) {
		t.Fatal(err)
	}
	if err := Drop(dir, "s"); err != nil {
		t.Fatal(err)
	}
	if m := current(t, dir); len(m.Sections) != 1 {
		t.Fatal(m.Sections)
	}
	file := filepath.Join(dir, "f")
	os.WriteFile(file, nil, 0o644)
	if err := storeGrammar(file, &Manifest{}, g); err == nil {
		t.Fatal("under a file")
	}
	if err := storeGrammar(file, &Manifest{Grammar: "set"}, g); err != nil {
		t.Fatal(err)
	}
	gd := filepath.Join(dir, dirDecl)
	os.RemoveAll(gd)
	os.WriteFile(gd, nil, 0o644)
	if err := Declare(dir, s, g); err == nil {
		t.Fatal("decl under a file")
	}
	bad := t.TempDir()
	os.WriteFile(filepath.Join(bad, currentFile), []byte("9"), 0o644)
	if err := Declare(bad, s, g); err == nil {
		t.Fatal("broken index")
	}
	if err := Drop(bad, "s"); err == nil {
		t.Fatal("broken index")
	}
	other := t.TempDir()
	os.WriteFile(filepath.Join(other, dirGrammar), nil, 0o644)
	if err := Declare(other, s, g); err == nil {
		t.Fatal("grammar under a file")
	}
}
