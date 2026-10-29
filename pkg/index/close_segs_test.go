//ff:func feature=index type=reader control=sequence
//ff:what TestCloseSegs — every file is closed (closing again fails)
package index

import "testing"

func TestCloseSegs(t *testing.T) {
	_, _, dir := built(t)
	m, _ := CurrentManifest(dir)
	ss, _ := openSegs(dir, m)
	closeSegs(ss)
	if ss[1][0].Close() == nil {
		t.Fatal("still open")
	}
}
