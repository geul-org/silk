//ff:func feature=query type=scanner control=sequence
//ff:what TestRegions — named partitions when present, else every partition
package query

import "testing"

func TestRegions(t *testing.T) {
	_, _, ix := smallIndex(t)
	f := ix.Section("label", 0).Files[0]
	k, _ := f.PartAt(0)
	if r := regions(&Plan{Parts: [][]byte{k, make([]byte, 16)}}, f); len(r) != 1 {
		t.Fatal(r)
	}
	if r := fixedRegions(&Plan{Parts: [][]byte{make([]byte, 16)}}, f); len(r) != 0 {
		t.Fatal(r)
	}
	if r := regions(&Plan{}, f); uint64(len(r)) != f.Parts {
		t.Fatal(r)
	}
}
