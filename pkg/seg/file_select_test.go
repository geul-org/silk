//ff:func feature=seg type=reader control=iteration dimension=1
//ff:what TestFileSelect — one partition, a key in the other partition, every partition with a sort key, and a read error in the partition loop
package seg

import (
	"fmt"
	"testing"
)

func TestFileSelect(t *testing.T) {
	cases := []struct {
		sel  Sel
		want []string
	}{
		{Sel{Part: []byte("aa"), Sort: []byte("x1")}, []string{"1:4", "2:4"}},
		{Sel{Part: []byte("cc")}, []string{"1:8", "4:2"}},
		{Sel{Part: []byte("bb"), Sort: []byte("y1")}, []string{"5:0"}},
		{Sel{Part: []byte("ee")}, nil},
		{Sel{Sort: []byte("w1")}, []string{"6:0"}},
		{Sel{Sort: []byte("z0")}, []string{"4:2"}},
	}
	s, u := sampleSeg(t, true), sampleSeg(t, false)
	for i, c := range cases {
		expectSelect(t, s, c.sel, c.want, fmt.Sprint("sorted ", i))
		expectSelect(t, u, c.sel, c.want, fmt.Sprint("appended ", i))
	}
	emit, _ := collect()
	if _, err := failing(u, 0).Select(Sel{}, emit); err == nil {
		t.Fatal("read error")
	}
}
