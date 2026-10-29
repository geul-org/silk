//ff:func feature=seg type=reader control=iteration dimension=1
//ff:what TestFileSelectRegion — binary search in a sorted run, a scan in an appended segment, and errors from either bound
package seg

import (
	"slices"
	"testing"
)

func TestFileSelectRegion(t *testing.T) {
	for _, sorted := range []bool{true, false} {
		f := sampleSeg(t, sorted)
		emit, got := collect()
		n, err := f.selectRegion(f.partRegion(0), []byte("x1"), 0, emit)
		slices.Sort(*got)
		if err != nil || !slices.Equal(*got, []string{"1:4", "2:4"}) || (sorted && n != 2) || (!sorted && n != 3) {
			t.Fatal(sorted, n, err, *got)
		}
	}
	f := sampleSeg(t, true)
	emit, _ := collect()
	if _, err := failing(f, 0).selectRegion(f.partRegion(0), []byte("x1"), 0, emit); err == nil {
		t.Fatal("lower bound")
	}
	if _, err := failing(f, 2).selectRegion(f.partRegion(0), []byte("x1"), 0, emit); err == nil {
		t.Fatal("upper bound")
	}
}
