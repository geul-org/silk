//ff:func feature=idle type=reader control=sequence
//ff:what Read — a Sample of this machine from /proc/loadavg and /proc/meminfo; ErrUnsupported without them
package idle

import "runtime"

// Read samples the system now.
func Read() (Sample, error) {
	return readFrom(loadavgPath, meminfoPath, runtime.NumCPU())
}
