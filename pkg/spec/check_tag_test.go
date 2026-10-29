//ff:func feature=spec type=rule control=iteration dimension=1
//ff:what TestCheckTag — a known layout with valid keys passes; unknown layout, missing name, negative minimum, bad partition or sort key fail
package spec

import "testing"

func TestCheckTag(t *testing.T) {
	ok := Tag{Name: "TRIPLE", Kind: KindEdge, Codec: 0x0A, Partition: Key{Path: []int{1}, Width: 6}, Sort: Key{Path: []int{3}, Width: 6}}
	if err := checkTag(ok); err != nil {
		t.Fatal(err)
	}
	bad := []func(*Tag){
		func(g *Tag) { g.Codec = 0x7E },
		func(g *Tag) { g.Name = "" },
		func(g *Tag) { g.MinPartition = -1 },
		func(g *Tag) { g.Partition = Key{Path: []int{3}, Width: 6} },
		func(g *Tag) { g.Sort = Key{Path: []int{9}, Width: 6} },
	}
	for i, f := range bad {
		g := ok
		f(&g)
		if checkTag(g) == nil {
			t.Fatalf("case %d passed", i)
		}
	}
}
