//ff:func feature=index type=writer control=sequence
//ff:what pinGen — mark a generation as in use by this process (pin/<generation>-<pid>-<nanotime>); the pin file path
package index

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// pinGen creates the pin; removing the file releases it.
func pinGen(dir string, gen uint64) (string, error) {
	pd := filepath.Join(dir, dirPin)
	if err := os.MkdirAll(pd, 0o755); err != nil {
		return "", err
	}
	p := filepath.Join(pd, fmt.Sprintf("%d-%d-%d", gen, os.Getpid(), time.Now().UnixNano()))
	return p, os.WriteFile(p, nil, 0o644)
}
