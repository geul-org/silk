//ff:func feature=seg type=helper control=sequence
//ff:what TestFilePartRegion — partition i's first row offset, count and row width
package seg

import "testing"

func TestFilePartRegion(t *testing.T) {
	f := sampleSeg(t, true)
	if r := f.partRegion(1); r.off != 30 || r.n != 2 || r.w != 10 {
		t.Fatal(r)
	}
}
