//ff:func feature=seg type=helper control=sequence
//ff:what TestAllSorted — true only when every group comes from a sorted run
package seg

import "testing"

func TestAllSorted(t *testing.T) {
	s, u := sampleSeg(t, true), sampleSeg(t, false)
	if !allSorted([]Group{{f: s}, {f: s}}) || allSorted([]Group{{f: s}, {f: u}}) {
		t.Fatal("allSorted")
	}
}
