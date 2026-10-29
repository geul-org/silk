//ff:func feature=seg type=reader control=sequence
//ff:what TestAdvance — a row, then nil at the end, and the read error
package seg

import "testing"

func TestAdvance(t *testing.T) {
	f := sampleSeg(t, true)
	c := &cursor{g: Group{f: f, rg: f.partRegion(1)}}
	a, _ := advance(c)
	advance(c)
	end, err := advance(c)
	if a == nil || end != nil || err != nil {
		t.Fatal(a, end, err)
	}
	if _, err := advance(&cursor{g: Group{f: failing(f, 0), rg: f.partRegion(1)}}); err == nil {
		t.Fatal("read error")
	}
}
