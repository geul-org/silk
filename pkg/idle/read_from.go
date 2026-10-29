//ff:func feature=idle type=reader control=sequence
//ff:what readFrom — a Sample from a loadavg file and a meminfo file and a CPU count
package idle

import (
	"fmt"
	"os"
)

// readFrom reads the two files (tests point it elsewhere).
func readFrom(loadPath, memPath string, cpus int) (Sample, error) {
	lb, err := os.ReadFile(loadPath)
	if err != nil {
		return Sample{}, fmt.Errorf("%w: %v", ErrUnsupported, err)
	}
	mb, err := os.ReadFile(memPath)
	if err != nil {
		return Sample{}, fmt.Errorf("%w: %v", ErrUnsupported, err)
	}
	l, err := parseLoadavg(lb)
	if err != nil {
		return Sample{}, err
	}
	t, a, err := parseMeminfo(mb)
	return Sample{Load1: l, CPUs: cpus, MemTotal: t, MemAvail: a}, err
}
