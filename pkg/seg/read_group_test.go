//ff:func feature=seg type=reader control=sequence
//ff:what TestReadGroup — all rows of a group (key stripped); a read error fails
package seg

import "testing"

func TestReadGroup(t *testing.T) {
	f := sampleSeg(t, true)
	gs, _ := f.Groups()
	if b, err := readGroup(gs[2]); err != nil || len(b) != 10 || string(b[:2]) != "y1" {
		t.Fatal(err, b)
	}
	if _, err := readGroup(Group{f: failing(f, 0), rg: gs[0].rg}); err == nil {
		t.Fatal("read error")
	}
}
