//ff:func feature=idle type=scanner control=sequence
//ff:what TestWait — returns when idle or unreadable, waits while busy, stops with the context
package idle

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestWait(t *testing.T) {
	r := Rule{CPUFree: .5, MemFree: .5}
	ctx, cancel := context.WithCancel(context.Background())
	n := 0
	sample := func() (Sample, error) {
		n++
		if n == 1 {
			return Sample{}, nil
		}
		return Sample{CPUs: 4, MemTotal: 10, MemAvail: 9}, nil
	}
	if err := Wait(ctx, r, time.Millisecond, sample); err != nil || n != 2 {
		t.Fatal(err, n)
	}
	if err := Wait(ctx, r, time.Hour, func() (Sample, error) { return Sample{}, errors.New("x") }); err != nil {
		t.Fatal(err)
	}
	go cancel()
	if err := Wait(ctx, r, time.Hour, func() (Sample, error) { return Sample{}, nil }); !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
	if err := Wait(ctx, r, time.Hour, sample); !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}
