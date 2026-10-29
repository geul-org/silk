//ff:func feature=cli type=helper control=sequence
//ff:what TestEdgeList — `#a,#b`, empty for none
package cli

import "testing"

func TestEdgeList(t *testing.T) {
	if edgeList([]uint64{7, 12}) != "#7,#12" || edgeList(nil) != "" {
		t.Fatal(edgeList([]uint64{7, 12}))
	}
}
