//ff:func feature=index type=writer control=sequence
//ff:what writeAtomic — put bytes at path through path.tmp (synced, then renamed), creating the directory; the bytes written
package index

import (
	"errors"
	"os"
	"path/filepath"
)

// writeAtomic never leaves a half-written file at path.
func writeAtomic(path string, b []byte) (int64, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return 0, err
	}
	f, err := os.Create(path + ".tmp")
	if err != nil {
		return 0, err
	}
	_, werr := f.Write(b)
	if err := errors.Join(werr, f.Sync(), f.Close()); err != nil {
		return 0, err
	}
	return int64(len(b)), os.Rename(path+".tmp", path)
}
