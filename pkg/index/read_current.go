//ff:func feature=index type=reader control=sequence
//ff:what readCurrent — the generation CURRENT names; ErrNoIndex when there is none
package index

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// readCurrent reads dir/CURRENT ("<generation>\n").
func readCurrent(dir string) (uint64, error) {
	b, err := os.ReadFile(filepath.Join(dir, currentFile))
	if os.IsNotExist(err) {
		return 0, fmt.Errorf("%w: %s", ErrNoIndex, dir)
	}
	if err != nil {
		return 0, err
	}
	g, err := strconv.ParseUint(strings.TrimSpace(string(b)), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: CURRENT %q", ErrFormat, b)
	}
	return g, nil
}
