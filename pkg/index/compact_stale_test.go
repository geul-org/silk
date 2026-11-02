//ff:func feature=index type=scanner control=sequence
//ff:what TestCompactStale — a section past StaleRatio drops its dead rows; a section of only dead rows ends empty
package index

import (
	"context"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/archive"
)

func TestCompactStale(t *testing.T) {
	ctx := context.Background()
	data, dir, a := built(t, 4)
	fixture.Repack(t, data, docsOf(1), nil)
	a, _ = archive.Open(ctx, data)
	o := Options{Limit: 1 << 21, Policy: DefaultPolicy()}
	if _, err := Update(ctx, a, dir, o); err != nil {
		t.Fatal(err)
	}
	if _, err := Compact(ctx, dir, o, false); err != nil {
		t.Fatal(err)
	}
	m := current(t, dir)
	if rowsOf(m.Sections[0].Runs) != 1 {
		t.Fatal(m.Sections[0])
	}
	c := &compactor{dir: dir, o: o, m: m, tmp: t.TempDir()}
	c.live = Bitmap{}
	if ts, err := c.plan(); err != nil || len(ts) == 0 {
		t.Fatal(ts, err)
	}
	s := &c.m.Sections[0]
	if err := c.merge(ctx, task{sec: s, runs: s.Runs}); err != nil {
		t.Fatal(err)
	}
	if len(s.Runs) != 0 {
		t.Fatal(s)
	}
}
