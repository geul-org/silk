//ff:func feature=grammar type=scanner control=sequence
//ff:what TestAddDocErrors — a published codec spec adds fields; an unreadable or undecodable version stops
package grammar

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestAddDocErrors(t *testing.T) {
	ctx := context.Background()
	src := filepath.Join(repoRoot(), "testdata", "grammar", "0x0000")
	a, err := pack.OpenDir(ctx, src, pack.BuildOptions{})
	if err != nil {
		t.Fatal(err)
	}
	g := &Grammar{Fields: map[uint16][]Field{}, Names: map[string][]uint16{}, Odd: map[uint16]bool{}, Langs: map[string]uint16{}}
	d := resolve.Dir{Store: a}
	if err := g.addDoc(d, "0x0000000000001004"); err != nil || len(g.Fields) != 1 {
		t.Fatal(err, len(g.Fields))
	}
	if err := g.addDoc(d, "0x0000000000000010"); err != nil {
		t.Fatal(err)
	}
	if err := g.addDoc(resolve.Dir{Store: brokenStore{a}}, "0x0000000000001004"); err == nil {
		t.Fatal("broken store")
	}
	if err := g.addDoc(resolve.Dir{Store: junkStore{a}}, "0x0000000000001004"); err == nil {
		t.Fatal("junk document")
	}
}
