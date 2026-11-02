//ff:func feature=seg type=writer control=sequence
//ff:what NewWriter — start a segment at path (written to path.tmp until Finish)
package seg

import (
	"bufio"
	"crypto/sha256"
	"os"
	"path/filepath"
)

func NewWriter(path string, info Info) (*Writer, error) {
	if err := info.check(); err != nil {
		return nil, err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	f, err := os.Create(path + ".tmp")
	if err != nil {
		return nil, err
	}
	tf, err := os.Create(path + ".tbl")
	if err != nil {
		f.Close()
		return nil, err
	}
	return &Writer{info: info, path: path, f: f, tf: tf, bw: bufio.NewWriterSize(f, bufSize), tw: bufio.NewWriterSize(tf, 1<<16), sum: sha256.New(), last: make([]byte, info.RowWidth)}, nil
}
