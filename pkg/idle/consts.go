// Package idle decides whether the machine is idle enough for background
// index maintenance. It starts simple: the 1-minute load average and the
// available memory from /proc (Linux). Work may start when the load plus the
// work's own threads leaves CPUFree of the CPUs unused and MemFree of the
// memory is available. Under WSL the load average covers only the Linux VM,
// not the Windows host.
package idle

import "errors"

const (
	// DefaultFree is the share of CPU and memory left for the user (30%).
	DefaultFree = 0.30
	loadavgPath = "/proc/loadavg"
	meminfoPath = "/proc/meminfo"
)

// ErrUnsupported means the system offers no load average or memory figures
// (no /proc): the caller decides (manual compaction still works).
var ErrUnsupported = errors.New("silk/idle: no /proc/loadavg or /proc/meminfo")
