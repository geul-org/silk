//ff:func feature=index type=writer control=sequence
//ff:what lockDir — take the index's writer lock (LOCK holding our process ID, created exclusively); a lock of a dead process is taken over; the release function
package index

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// lockDir allows one writer (build, update, compact) at a time. Readers do
// not lock; they pin a generation.
func lockDir(dir string) (func(), error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	p := filepath.Join(dir, lockFile)
	f, err := os.OpenFile(p, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
	if os.IsExist(err) {
		b, _ := os.ReadFile(p)
		if pid, _ := strconv.Atoi(strings.TrimSpace(string(b))); pidAlive(pid) {
			return nil, fmt.Errorf("%w (process %d)", ErrLocked, pid)
		}
		os.Remove(p)
		f, err = os.OpenFile(p, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
	}
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(f, "%d\n", os.Getpid())
	f.Close()
	return func() { os.Remove(p) }, nil
}
