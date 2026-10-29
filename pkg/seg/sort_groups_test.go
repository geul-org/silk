//ff:func feature=seg type=helper control=sequence
//ff:what TestSortGroups — groups ordered by key, equal keys keep their order
package seg

import "testing"

func TestSortGroups(t *testing.T) {
	gs := []Group{{Key: []byte("b"), strip: 1}, {Key: []byte("a")}, {Key: []byte("b"), strip: 2}}
	SortGroups(gs)
	if string(gs[0].Key) != "a" || gs[1].strip != 1 || gs[2].strip != 2 {
		t.Fatal(gs)
	}
}
