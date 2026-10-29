//go:build unix

//ff:func feature=index type=helper control=sequence
//ff:what pidAlive — whether a process with this ID exists (signal 0; permission denied still means it exists)
package index

import (
	"errors"
	"syscall"
)

// pidAlive tells stale locks and pins from live ones.
func pidAlive(pid int) bool {
	if pid <= 0 {
		return false
	}
	err := syscall.Kill(pid, 0)
	return err == nil || errors.Is(err, syscall.EPERM)
}
