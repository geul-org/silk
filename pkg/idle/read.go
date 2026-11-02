//ff:func feature=idle type=reader control=sequence
//ff:what Read — a Sample of this machine; ErrUnsupported without /proc
package idle

import "runtime"

func Read() (Sample, error) {
	return readFrom(loadavgPath, meminfoPath, runtime.NumCPU())
}
