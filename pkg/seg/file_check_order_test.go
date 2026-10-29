//ff:func feature=seg type=rule control=sequence
//ff:what TestFileCheckOrder — ordered rows pass; descending prefixes, descending rows in a sorted run and read errors fail
package seg

import "testing"

func TestFileCheckOrder(t *testing.T) {
	f := sampleSeg(t, true)
	if err := f.checkOrder(Group{f: f, rg: f.otherRegion()}, 2, true); err != nil {
		t.Fatal(err)
	}
	u := sampleSeg(t, false)
	if err := u.checkOrder(Group{f: u, rg: u.partRegion(0)}, 0, false); err != nil {
		t.Fatal(err)
	}
	if err := u.checkOrder(Group{f: u, rg: u.partRegion(0)}, 0, true); err == nil {
		t.Fatal("unsorted rows in a sorted run")
	}
	if err := u.checkOrder(Group{f: u, rg: u.partRegion(0)}, 2, false); err == nil {
		t.Fatal("descending prefix")
	}
	if err := f.checkOrder(Group{f: failing(f, 0), rg: f.otherRegion()}, 2, true); err == nil {
		t.Fatal("read error")
	}
}
