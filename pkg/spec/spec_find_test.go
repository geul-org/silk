//ff:func feature=spec type=helper control=sequence
//ff:what TestSpecFind — the tag of an indexed TAG WORD, false for another
package spec

import "testing"

func TestSpecFind(t *testing.T) {
	if g, ok := Default().Find(0x100B); !ok || g.Name != "QUALIFIER" {
		t.Fatal(g)
	}
	if _, ok := Default().Find(0x0801); ok {
		t.Fatal("found a Value kind")
	}
}
