//ff:func feature=query type=scanner control=iteration dimension=1
//ff:what TestSelectCond — candidates sorted by document and offset with the lookup counted; a failing segment read fails
package query

import (
	"testing"
)

func TestSelectCond(t *testing.T) {
	ix := testIndex(t, 1)
	p, _ := planCond(ix, condOf(t, "P106=*"))
	var st Stats
	cs, err := selectCond(ix, p, &st)
	if err != nil || len(cs) != 7 || st.Candidates != 7 || st.Segments != 1 || st.Rows != 7 {
		t.Fatal(err, st)
	}
	for i := 1; i < len(cs); i++ {
		if cs[i].doc < cs[i-1].doc || (cs[i].doc == cs[i-1].doc && cs[i].off <= cs[i-1].off) {
			t.Fatal(cs)
		}
	}
	ix.Segs[p.Section][0].Close()
	if _, err := selectCond(ix, p, &st); err == nil {
		t.Fatal("closed segment")
	}
}
