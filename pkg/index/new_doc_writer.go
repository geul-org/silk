//ff:func feature=index type=writer control=sequence
//ff:what newDocWriter — start a document table file at path
package index

import (
	"bufio"
	"crypto/sha256"
	"os"
	"path/filepath"
)

func newDocWriter(path string, first uint32) (*docWriter, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	f, err := os.Create(path + ".tmp")
	if err != nil {
		return nil, err
	}
	return &docWriter{path: path, f: f, bw: bufio.NewWriterSize(f, 1<<20), sum: sha256.New(), first: first}, nil
}
