//ff:func feature=grammar type=scanner control=iteration dimension=1
//ff:what TestGrammarAddDoc — a document that is no codec spec is skipped; an unpublished one adds nothing
package grammar

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestGrammarAddDoc(t *testing.T) {
	g := &Grammar{}
	if err := g.addDoc(resolve.Dir{Root: t.TempDir()}, "versions.geul"); err != nil {
		t.Fatal(err)
	}
	if err := g.addDoc(resolve.Dir{Root: t.TempDir()}, "0x0000000000000848"); err != nil {
		t.Fatal(err)
	}
	if id, ok := parseDocID("0x0000000000000848"); !ok || len(id) != 8 {
		t.Fatal(id, ok)
	}
	for _, n := range []string{"0x0000000000000000", "0xffff000000000002", "0x0000000000000888", "0xzz00000000000848", "abc"} {
		if _, ok := parseDocID(n); ok {
			t.Fatal(n)
		}
	}
}
