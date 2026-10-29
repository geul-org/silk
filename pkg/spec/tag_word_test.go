//ff:func feature=spec type=helper control=sequence
//ff:what TestTagWord — TRIPLE is 0x100A (as gogeul core.TargetWord), extra bits are masked
package spec

import "testing"

func TestTagWord(t *testing.T) {
	if TagWord(KindEdge, 0x0A) != 0x100A || TagWord(1, 0x4A) != 0x084A || TagWord(6, 0xFF) != 0x107F {
		t.Fatal(TagWord(KindEdge, 0x0A), TagWord(1, 0x4A), TagWord(6, 0xFF))
	}
}
