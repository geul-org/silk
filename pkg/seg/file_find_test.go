//ff:func feature=seg type=reader control=iteration dimension=1
//ff:what TestFileFind — table keys are found by binary search; keys in the other partition or absent are not
package seg

import "testing"

func TestFileFind(t *testing.T) {
	f := sampleSeg(t, true)
	if i, ok := f.Find([]byte("cc")); !ok || i != 1 {
		t.Fatal(i, ok)
	}
	for _, k := range []string{"bb", "zz", "00"} {
		if _, ok := f.Find([]byte(k)); ok {
			t.Fatal(k)
		}
	}
}
