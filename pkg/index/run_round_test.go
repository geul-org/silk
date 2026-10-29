//ff:func feature=index type=builder control=iteration dimension=1
//ff:what TestRunRound — every task of a round runs (one thread or the default), outcomes in task order
package index

import (
	"context"
	"testing"
)

func TestRunRound(t *testing.T) {
	_, _, dir := updated(t)
	m, _ := CurrentManifest(dir)
	live, _ := readBitmap(dir + "/" + m.Live)
	for _, th := range []int{1, 0} {
		ts := planNamed(&m, smallSpec(), Options{Full: true})
		outs := runRound(context.Background(), dir, ts, nil, live, Options{Threads: th})
		if len(outs) != 7 || outs[3].err != nil || outs[3].res.Task.Section != 3 || outs[3].res.Out.Rows == 0 {
			t.Fatal(th, outs[3])
		}
	}
}
