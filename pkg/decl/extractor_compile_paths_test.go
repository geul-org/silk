//ff:func feature=decl type=scanner control=sequence
//ff:what TestExtractorCompilePaths — partition and keys resolve on the first TAG
package decl

import (
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestExtractorCompilePaths(t *testing.T) {
	g := fixture.Grammar(t)
	x := &Extractor{g: g, d: Decl{Tags: []uint16{0x1004}, Part: "object.language", Keys: []string{"subject"}}}
	if err := x.compilePaths(); err != nil || x.part == nil || len(x.keys) != 1 {
		t.Fatal(err)
	}
	x = &Extractor{g: g, d: Decl{Tags: []uint16{0x1004}, Part: "nope"}}
	if err := x.compilePaths(); err == nil {
		t.Fatal("part")
	}
	x = &Extractor{g: g, d: Decl{Tags: []uint16{0x1004}, Keys: []string{"nope"}}}
	if err := x.compilePaths(); err == nil {
		t.Fatal("key")
	}
}
