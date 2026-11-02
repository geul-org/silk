//ff:func feature=idle type=scanner control=sequence
//ff:what TestRule — the rule must be set; idle needs both shares free
package idle

import (
	"errors"
	"testing"
)

func TestRule(t *testing.T) {
	if err := (Rule{}).Check(); !errors.Is(err, ErrRule) {
		t.Fatal(err)
	}
	r := Rule{CPUFree: .5, MemFree: .5, Threads: 1}
	if err := r.Check(); err != nil {
		t.Fatal(err)
	}
	if err := (Rule{CPUFree: 1, MemFree: .5}).Check(); !errors.Is(err, ErrRule) {
		t.Fatal(err)
	}
	s := Sample{Load1: 0, CPUs: 4, MemTotal: 100, MemAvail: 60}
	if !r.Idle(s) {
		t.Fatal("idle")
	}
	if r.Idle(Sample{Load1: 3.5, CPUs: 4, MemTotal: 100, MemAvail: 60}) || r.Idle(Sample{CPUs: 4, MemTotal: 100, MemAvail: 10}) || r.Idle(Sample{}) || r.Idle(Sample{CPUs: 1}) {
		t.Fatal("busy")
	}
}
