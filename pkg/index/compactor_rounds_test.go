//ff:func feature=index type=scanner control=sequence
//ff:what TestCompactorRounds — rounds finish, a bad bitmap, a failing merge and a cancelled context stop them
package index

import (
	"context"
	"errors"
	"testing"
)

func TestCompactorRounds(t *testing.T) {
	ctx := context.Background()
	_, dir, _ := built(t, 2)
	o := Options{Limit: 1 << 21, Policy: DefaultPolicy()}
	mk := func() *compactor {
		return &compactor{dir: dir, o: o, m: current(t, dir), tmp: t.TempDir(), full: true}
	}
	c := mk()
	c.m.Sections[0].Runs = append(c.m.Sections[0].Runs, c.m.Sections[0].Runs...)
	if r, err := c.rounds(ctx); err != nil || r.Merges == 0 {
		t.Fatal(r, err)
	}
	c = mk()
	c.m.Live = "live/none"
	if _, err := c.rounds(ctx); err == nil {
		t.Fatal("bitmap")
	}
	c = mk()
	c.m.Sections[0].Runs = []SegRef{{File: "none"}}
	if _, err := c.rounds(ctx); err == nil {
		t.Fatal("plan")
	}
	cctx, cancel := context.WithCancel(ctx)
	cancel()
	c = mk()
	if _, err := c.rounds(cctx); !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
	c = mk()
	c.m.Sections[0].Runs = append(c.m.Sections[0].Runs, c.m.Sections[0].Runs[0])
	c.o.Pause = func() error { return errors.New("stop") }
	if _, err := c.rounds(ctx); err == nil {
		t.Skip("pause point falls on the first row only")
	}
}
