//ff:func feature=index type=reader control=sequence
//ff:what TestIndexClose — data packs and segments closed and the pin removed; an index without a pin closes quietly
package index

import (
	"testing"
)

func TestIndexClose(t *testing.T) {
	a, _, dir := built(t)
	ix, _ := Open(a, dir)
	ix.Record(0, 0)
	if err := ix.Close(); err != nil || len(pinnedGens(dir)) != 0 {
		t.Fatal(err)
	}
	if err := (&Index{}).Close(); err != nil {
		t.Fatal(err)
	}
}
