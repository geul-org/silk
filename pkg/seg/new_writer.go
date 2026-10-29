//ff:func feature=seg type=writer control=sequence
//ff:what NewWriter — start a segment at path (written to path.tmp until Finish) with the given TAG, key widths, sorted flag and spec hash
package seg

import (
	"bufio"
	"crypto/sha256"
	"os"
	"path/filepath"
)

// NewWriter creates the temporary file; info's counts and Sum are ignored.
func NewWriter(path string, info Info, min int) (*Writer, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	f, err := os.OpenFile(path+".tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, err
	}
	info.Parts, info.Rows, info.Other, info.Sum, info.Bytes = 0, 0, 0, [32]byte{}, 0
	return &Writer{Path: path, Min: min, info: info, f: f, bw: bufio.NewWriterSize(f, 1<<20), h: sha256.New()}, nil
}
