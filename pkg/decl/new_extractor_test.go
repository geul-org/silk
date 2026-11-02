//ff:func feature=decl type=scanner control=sequence
//ff:what TestNewExtractor — paths that do not resolve and definitions that do not parse fail
package decl

import (
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestNewExtractor(t *testing.T) {
	g := fixture.Grammar(t)
	if _, err := NewExtractor(Decl{Kind: KindIndex, Tags: []uint16{0x1004}, Part: "nope"}, g); err == nil {
		t.Fatal("part")
	}
	if _, err := NewExtractor(Decl{Kind: KindIndex, Tags: []uint16{0x1004}, Keys: []string{"nope"}}, g); err == nil {
		t.Fatal("key")
	}
	if _, err := NewExtractor(Decl{Kind: KindSilk, Silk: "x"}, g); err == nil {
		t.Fatal("silk")
	}
}
