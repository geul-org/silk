//ff:func feature=query type=scanner control=sequence
//ff:what TestRunCond — documents meeting one condition, prefix candidates that fail the check dropped; lookup and check errors fail
package query

import (
	"testing"
)

func TestRunCond(t *testing.T) {
	ix := testIndex(t, 256)
	var st Stats
	p, _ := planCond(ix, condOf(t, "REFERENCE-NODE:1=#189"))
	hs, err := runCond(ix, p, 2, &st)
	if err != nil || len(hs) != 1 || hs[0].edges[0] != 188 || st.Verified != 1 || st.Candidates <= 1 {
		t.Fatal(err, hs, st)
	}
	bad := p
	bad.Cond.Codec = 0x0A
	if _, err := runCond(ix, bad, 2, &st); err == nil {
		t.Fatal("check")
	}
	ix.Segs[p.Section][0].Close()
	if _, err := runCond(ix, p, 2, &st); err == nil {
		t.Fatal("lookup")
	}
}
