//ff:func feature=idle type=reader control=sequence
//ff:what readFrom — a Sample from these two files
package idle

import (
	"os"
	"strconv"
	"strings"
)

func readFrom(loadavg, meminfo string, cpus int) (Sample, error) {
	l, err1 := os.ReadFile(loadavg)
	m, err2 := os.ReadFile(meminfo)
	if err1 != nil || err2 != nil {
		return Sample{}, ErrUnsupported
	}
	f := strings.Fields(string(l))
	if len(f) == 0 {
		return Sample{}, ErrUnsupported
	}
	load, err := strconv.ParseFloat(f[0], 64)
	if err != nil {
		return Sample{}, ErrUnsupported
	}
	s := Sample{Load1: load, CPUs: cpus}
	s.MemTotal, s.MemAvail = parseMeminfo(string(m))
	return s, nil
}
