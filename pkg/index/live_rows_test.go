//ff:func feature=index type=reader control=sequence
//ff:what TestLiveRows — rows of live documents with their partition slot; rows of dead ones counted
package index

import (
	"testing"
)

func TestLiveRows(t *testing.T) {
	a, _, dir := built(t)
	ix, _ := Open(a, dir)
	defer ix.Close()
	all, dead, err := liveRows(ix.Segs[0][0], ix.Live, nil)
	if err != nil || dead != 0 || len(all)%20 != 0 {
		t.Fatal(err, dead)
	}
	ix.Live.Put(0, false)
	some, dead2, _ := liveRows(ix.Segs[0][0], ix.Live, nil)
	if dead2 == 0 || uint64(len(some)/20)+dead2 != uint64(len(all)/20) {
		t.Fatal(dead2)
	}
}
