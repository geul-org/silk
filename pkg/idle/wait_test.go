//ff:func feature=idle type=scanner control=sequence
//ff:what TestWait — returns at once when idle or unreadable, polls while busy until idle, and stops with the context
package idle

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestWait(t *testing.T) {
	ctx := context.Background()
	busy := Sample{Load1: 99, CPUs: 4, MemTotal: 1, MemAvail: 1}
	quiet := Sample{Load1: 0, CPUs: 4, MemTotal: 1, MemAvail: 1}
	if err := Wait(ctx, Rule{}, time.Millisecond, func() (Sample, error) { return Sample{}, ErrUnsupported }); err != nil {
		t.Fatal(err)
	}
	n := 0
	if err := Wait(ctx, Rule{}, time.Millisecond, func() (Sample, error) {
		n++
		if n < 3 {
			return busy, nil
		}
		return quiet, nil
	}); err != nil || n != 3 {
		t.Fatal(err, n)
	}
	c1, cancel := context.WithTimeout(ctx, 5*time.Millisecond)
	defer cancel()
	if err := Wait(c1, Rule{}, time.Hour, func() (Sample, error) { return busy, nil }); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatal(err)
	}
	c2, cancel2 := context.WithCancel(ctx)
	cancel2()
	if err := Wait(c2, Rule{}, time.Hour, func() (Sample, error) { return quiet, nil }); !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}
