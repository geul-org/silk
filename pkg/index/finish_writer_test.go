//ff:func feature=index type=writer control=sequence
//ff:what TestFinishWriter — without compaction only the time is set; with it the merges are added to the report
package index

import (
	"context"
	"testing"
	"time"
)

func TestFinishWriter(t *testing.T) {
	_, _, dir := updated(t)
	m, _ := CurrentManifest(dir)
	t0 := time.Now()
	r, err := finishWriter(context.Background(), dir, &m, Options{}, Report{Written: 5}, t0)
	if err != nil || r.Written != 5 || len(r.Tasks) != 0 || r.Elapsed == 0 {
		t.Fatal(err, r)
	}
	r, err = finishWriter(context.Background(), dir, &m, Options{Compact: true, Full: true}, Report{Written: 5}, t0)
	if err != nil || r.Written <= 5 || len(r.Tasks) != 7 || r.Generation != m.Generation {
		t.Fatal(err, r)
	}
}
