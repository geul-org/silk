//ff:func feature=index type=scanner control=sequence
//ff:what TestCloseSegs — closing empties nothing it should not
package index

import "testing"

func TestCloseSegs(t *testing.T) {
	_, dir, _ := built(t, 1)
	fs, err := openSegs(segPaths(dir, current(t, dir).Sections[0]))
	if err != nil || len(fs) != 1 {
		t.Fatal(err)
	}
	closeSegs(fs)
	closeSegs(nil)
}
