//ff:func feature=seg type=reader control=sequence
//ff:what TestGroupsOf — groups of two segments interleaved by key; a failing read fails
package seg

import "testing"

func TestGroupsOf(t *testing.T) {
	a, b := sampleSeg(t, true), sampleSeg(t, false)
	gs, err := GroupsOf([]*File{a, b})
	if err != nil || len(gs) != 8 || string(gs[1].Key) != "aa" || gs[1].f != b {
		t.Fatal(err, len(gs))
	}
	if _, err := GroupsOf([]*File{failing(a, 0)}); err == nil {
		t.Fatal("read error")
	}
}
