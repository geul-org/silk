//ff:func feature=query type=matcher control=iteration dimension=1
//ff:what TestCheckAll — every candidate checked on several goroutines (or one); a candidate that cannot be read fails
package query

import (
	"testing"
)

func TestCheckAll(t *testing.T) {
	ix := testIndex(t, 1)
	c := condOf(t, "P106=*")
	p, _ := planCond(ix, c)
	var st Stats
	cs, _ := selectCond(ix, p, &st)
	for _, th := range []int{4, 0} {
		ids, ok, err := checkAll(ix, cs, c, th)
		if err != nil || len(ids) != 7 || !ok[6] || ids[0] == 0 {
			t.Fatal(err, ids, ok)
		}
	}
	if _, _, err := checkAll(ix, []candidate{{doc: 99}}, c, 2); err == nil {
		t.Fatal("unknown document")
	}
}
