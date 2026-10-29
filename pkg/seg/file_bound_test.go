//ff:func feature=seg type=reader control=sequence
//ff:what TestFileBound — lower and upper bounds of a prefix in a sorted partition; a failing read is reported
package seg

import "testing"

func TestFileBound(t *testing.T) {
	f := sampleSeg(t, true)
	rg := f.partRegion(0)
	lo, _ := f.bound(rg, []byte("x1"), false)
	hi, _ := f.bound(rg, []byte("x1"), true)
	end, _ := f.bound(rg, []byte("x9"), false)
	if lo != 0 || hi != 2 || end != 3 {
		t.Fatal(lo, hi, end)
	}
	if _, err := failing(f, 0).bound(rg, []byte("x1"), false); err == nil {
		t.Fatal("read error")
	}
}
