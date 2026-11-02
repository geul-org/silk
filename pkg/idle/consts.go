// Package idle decides whether the machine is idle enough for background
// merges, by a rule the caller sets: the share of CPU and memory that must
// stay free. There are no defaults. It reads the 1-minute load average and the
// available memory from /proc (Linux). Under WSL the load average covers only
// the Linux VM, not the Windows host.
package idle

import "errors"

const (
	loadavgPath = "/proc/loadavg"
	meminfoPath = "/proc/meminfo"
)

var (
	// ErrUnsupported means the system offers no load average or memory figures.
	ErrUnsupported = errors.New("silk/idle: no /proc/loadavg or /proc/meminfo")
	// ErrRule means the caller did not set the rule.
	ErrRule = errors.New("silk/idle: the free CPU and memory shares must be set")
)
