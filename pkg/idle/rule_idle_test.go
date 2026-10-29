//ff:func feature=idle type=rule control=sequence
//ff:what TestRuleIdle — defaults leave 30%: idle on a quiet machine, busy under load or low memory, never idle without figures
package idle

import "testing"

func TestRuleIdle(t *testing.T) {
	quiet := Sample{Load1: 1, CPUs: 28, MemTotal: 100, MemAvail: 50}
	if !(Rule{}).Idle(quiet) || !(Rule{CPUFree: 0.3, MemFree: 0.3, Threads: 2}).Idle(quiet) {
		t.Fatal("quiet machine")
	}
	if (Rule{}).Idle(Sample{Load1: 20, CPUs: 28, MemTotal: 100, MemAvail: 50}) {
		t.Fatal("loaded")
	}
	if (Rule{}).Idle(Sample{Load1: 0, CPUs: 28, MemTotal: 100, MemAvail: 20}) {
		t.Fatal("low memory")
	}
	if (Rule{}).Idle(Sample{}) {
		t.Fatal("no figures")
	}
}
