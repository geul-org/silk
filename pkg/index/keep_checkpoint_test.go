//ff:func feature=index type=helper control=sequence
//ff:what TestKeepCheckpoint — each paused output's temporary file joins the keep set
package index

import "testing"

func TestKeepCheckpoint(t *testing.T) {
	d := t.TempDir()
	writeCheckpoints(d, []checkpoint{{Output: "seg/a"}, {Output: "seg/b"}})
	k := map[string]bool{}
	keepCheckpoint(d, k)
	if !k["seg/a.tmp"] || !k["seg/b.tmp"] || len(k) != 2 {
		t.Fatal(k)
	}
}
