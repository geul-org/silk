//ff:func feature=index type=scanner control=sequence
//ff:what TestIndexLoadSections — built sections open, pending ones only declare, a missing segment or declaration fails
package index

import (
	"testing"

	"github.com/geul-org/silk/pkg/decl"
)

func TestIndexLoadSections(t *testing.T) {
	_, dir, _ := built(t, 1)
	m := current(t, dir)
	ix := &Index{Dir: dir, Manifest: m, Decls: map[string]decl.Decl{}, secs: map[string][]*Sec{}}
	if err := ix.loadSections(); err != nil || len(ix.secs) != 2 {
		t.Fatal(err)
	}
	ix.Close()
	p := Index{Dir: dir, Manifest: Manifest{Sections: []Section{{Name: "x", Decl: m.Sections[0].Decl, Pending: true}}}, Decls: map[string]decl.Decl{}, secs: map[string][]*Sec{}}
	if err := p.loadSections(); err != nil || p.secs["x"] != nil {
		t.Fatal(err)
	}
	b := Index{Dir: dir, Manifest: Manifest{Sections: []Section{{Decl: "none"}}}, Decls: map[string]decl.Decl{}}
	if err := b.loadSections(); err == nil {
		t.Fatal("no declaration")
	}
	c := Index{Dir: dir, Manifest: Manifest{Sections: []Section{{Decl: m.Sections[0].Decl, Runs: []SegRef{{File: "none"}}}}}, Decls: map[string]decl.Decl{}, secs: map[string][]*Sec{}}
	if err := c.loadSections(); err == nil {
		t.Fatal("no segment")
	}
}
