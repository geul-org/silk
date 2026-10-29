//ff:func feature=spec type=encoder control=sequence
//ff:what TestTagHash — a tag's hash is stable and changes with its key widths
package spec

import "testing"

func TestTagHash(t *testing.T) {
	a := Default().Tags[0]
	b := a
	b.Sort.Width = 8
	if a.Hash() != Default().Tags[0].Hash() || a.Hash() == b.Hash() {
		t.Fatal("hash")
	}
}
