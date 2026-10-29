//ff:func feature=seg type=helper control=sequence
//ff:what TestFileOtherRegion — the other partition starts after the partition rows with the wider rows
package seg

import "testing"

func TestFileOtherRegion(t *testing.T) {
	f := sampleSeg(t, true)
	if r := f.otherRegion(); r.off != 50 || r.n != 2 || r.w != 12 {
		t.Fatal(r)
	}
}
