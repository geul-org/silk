//ff:func feature=spec type=encoder control=sequence
//ff:what TestSpecHash — 64 hex digits, stable, and changed by any tag change
package spec

import "testing"

func TestSpecHash(t *testing.T) {
	s := Default()
	h := s.Hash()
	s.Tags[0].MinPartition++
	if len(h) != 64 || h != Default().Hash() || s.Hash() == h {
		t.Fatal(h)
	}
}
