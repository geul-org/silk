//ff:func feature=seg type=reader control=sequence
//ff:what TestFileGroups — table partitions then other-partition keys as groups; a failing read of the other partition fails
package seg

import "testing"

func TestFileGroups(t *testing.T) {
	f := sampleSeg(t, true)
	gs, err := f.Groups()
	if err != nil || len(gs) != 4 || string(gs[0].Key) != "aa" || string(gs[2].Key) != "bb" || gs[2].strip != 2 {
		t.Fatal(err, gs)
	}
	if _, err := failing(f, 0).Groups(); err == nil {
		t.Fatal("read error")
	}
}
